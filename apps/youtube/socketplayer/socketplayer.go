package socketplayer

import (
  "net/http"
  "github.com/googollee/go-socket.io"
  "fmt"
)

func SocketServe() {
    server, err := socketio.NewServer(nil)
    if err != nil {
        fmt.Println("[Fatal]",err)
    }
    server.On("connection", func(so socketio.Socket) {
        so.Join("player")


        so.On("play", func(msg string) {
            fmt.Println("emit:", so.Emit("play", msg))
            // so.BroadcastTo("chat", "chat message", msg)
        })
        so.On("pause", func(msg string) {
            // so.BroadcastTo("chat", "chat message", msg)
        })
        so.On("volume-change", func(msg string) {
            // so.BroadcastTo("chat", "chat message", msg)
        })

        so.On("response", func(msg string) {
            // so.BroadcastTo("chat", "chat message", msg)
        })


        so.On("disconnection", func() {
            fmt.Println("on disconnect")
        })
    })
    server.On("error", func(so socketio.Socket, err error) {
        fmt.Println("error:", err)
    })


    socketserver := http.NewServeMux()
    socketserver.Handle("/socket.io/", server)
    socketserver.Handle("/", http.FileServer(http.Dir("./asset")))

    fmt.Println(http.ListenAndServe(":5000", socketserver))
    fmt.Println("Serving at localhost:5000")
}
