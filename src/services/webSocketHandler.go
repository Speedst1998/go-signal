package services

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	localWebsocket "example.com/accounting/src/services/websocket"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func WebSocketHandler(c *gin.Context, MediaServerSockets map[string]localWebsocket.MediaServer, mediaServerName string)  {
    // Upgrade the HTTP connection to a WebSocket connection
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        http.Error(c.Writer, "Failed to upgrade to WebSocket", http.StatusBadRequest)
        return
    }

    MediaServerSockets[mediaServerName] = localWebsocket.MediaServer{Connection: conn, Name: mediaServerName}
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

func ClientConnect(c *gin.Context, MediaServerSockets map[string]localWebsocket.MediaServer, mediaServerName string, description string) (string, error)  {

    println("in clineconnect")
    mediaServer, ok := MediaServerSockets[mediaServerName]

    if ok {
        mediaServer.Connection.WriteMessage(websocket.TextMessage, []byte(description))
        println("WHATIS THE ERROR")

        _, message, _ := mediaServer.Connection.ReadMessage()
        fmt.Printf("Received message: %s\n", message)
        return string(message), nil
    }
    return "", errors.New("couldnt find media server")
}
