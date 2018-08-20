package main

import (
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)


type room struct {
    // forward holds incoming messages that should be forwardd to all other
    // connected clients.
    forward chan []byte

    // join is a channel for clients wishing to join the room.
    join chan *client

    // leave is a channel for clients wishing to leave the channel.
    leave chan *client

    // clients holds all clients currently connected to the room.
    clients map[*client]bool
}


func newRoom() *room {
    return &room{
        forward: make(chan []byte),
        join: make(chan *client),
        leave: make(chan *client),
        clients: make(map[*client]bool),
    }
}


func (r *room) run() {
    for {
        select {
        case client := <-r.join:
            // Joining room.
            r.clients[client] = true
        case client := <-r.leave:
            // Leaving room.
            delete(r.clients, client)
            close(client.send)
        case msg := <-r.forward:
            // Broadcast message.
            for client := range r.clients {
                client.send <- msg
            }
        }
    }
}


const (
    socketBufferSize = 1024
    messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
    ReadBufferSize: socketBufferSize,
    WriteBufferSize: socketBufferSize,
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    socket, err := upgrader.Upgrade(w, req, nil)
    if err != nil {
        log.Fatal("room ServeHTTP:", err)
        return
    }

    client := &client{
        socket: socket,
        send: make(chan []byte, messageBufferSize),
        room: r,
    }
    r.join <- client
    defer func() { r.leave <-client }()

    go client.write()
    client.read()
}
