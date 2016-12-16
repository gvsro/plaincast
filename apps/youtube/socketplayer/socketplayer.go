package socketplayer

import (
  "net/http"
	"github.com/googollee/go-socket.io"
  "github.com/gvsro/plaincast/log"
)

func main() {
    server, err := socketio.NewServer(nil)
    if err != nil {
        logger.Fatal(err)
    }
    server.On("connection", func(so socketio.Socket) {
        so.Join("player")


        so.On("play", func(msg string) {
            logger.Println("emit:", so.Emit("play", msg))
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
            logger.Println("on disconnect")
        })
    })
    server.On("error", func(so socketio.Socket, err error) {
        logger.Println("error:", err)
    })


    socketserver := http.NewServeMux()
    socketserver.Handle("/socket.io/", server)
    socketserver.Handle("/", http.FileServer(http.Dir("./asset")))


    logger.Println("Serving at localhost:5000...")
    logger.Println(http.ListenAndServe(":5000", socketserver))
}
