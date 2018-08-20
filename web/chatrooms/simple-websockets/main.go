package main

import (
    "flag"
    "log"
    "net/http"
    "path/filepath"
    "os"
    "sync"
    "text/template"

    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
)


type templateHandler struct {
    once sync.Once
    filename string
    template *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    t.once.Do(func() {
        t.template = template.Must(
            template.ParseFiles(filepath.Join("templates", t.filename)),
        )
    })

    t.template.Execute(w, nil)
}


var port string

func main() {
    flag.StringVar(&port, "port", "8080", "Port for server to connect to.")
    flag.Parse()

    //mux := http.NewServeMux()
    mux := mux.NewRouter()

    mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
    mux.Handle("/", &templateHandler{filename: "chat.html"})

    r := newRoom()
    mux.Handle("/room", r)

    // Get the room going.
    go r.run()

    address := "0.0.0.0:" + port
    log.Printf("Starting web server on: %s\n", address)
    err := http.ListenAndServe(address, handlers.LoggingHandler(os.Stdout, mux))
    if err != nil {
        log.Fatal(err)
    }
}
