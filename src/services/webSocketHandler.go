package services

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func WebSocketHandler(c *gin.Context, MediaServerSockets map[string]*websocket.Conn, mediaServerName string)  {
    // Upgrade the HTTP connection to a WebSocket connection
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        http.Error(c.Writer, "Failed to upgrade to WebSocket", http.StatusBadRequest)
        return
    }

    MediaServerSockets[mediaServerName] = conn
    // Loop to read messages from the WebSocket connection
    // for {
    //     _, message, err := conn.ReadMessage()
    //     if err != nil {
    //         break
    //     }
    //     // Handle the incoming message
    //     fmt.Printf("Received message: %s\n", message)
    // }
}

func ClientConnect(c *gin.Context, MediaServerSockets map[string]*websocket.Conn, mediaServerName string, description string) (string, error)  {

    println("in clineconnect")
    socket, ok := MediaServerSockets[mediaServerName]
    println(socket)
    if ok {
        socket.WriteMessage(websocket.TextMessage, []byte(description))
        println("WHATIS THE ERROR")

        _, message, _ := socket.ReadMessage()
        fmt.Printf("Received message: %s\n", message)
        return string(message), nil
    }
    return "", errors.New("couldnt find media server")
}
