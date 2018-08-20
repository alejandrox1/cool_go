function chatroom() {
    var socket = null;
    var msgBox = $("#chatbox textarea");
    var messages = $("#messages");

    alert("Doing stuff");

    $("#chatbox").submit(function() {
        if (!msgBox.val()) return false;
        if (!socket) {
            alert("Error: there is no socket connection.")
            return false;
        }

        socket.send(msgBox.val());
        msgBox.val("");
        return false;
    });

    if (!window["WebSocket"]) {
        alert("Error: your browswer doesn't seem to support websocket connections.")
    } else {
        socket = new WebSocket("ws://{{.Host}}/room");
        
        socket.onclose = function() {
            alert("Connection has been closed.");
        }

        socket.onmessage = function(e) {
            messages.append($("<li>").text(e.data));
        }
    }
}
