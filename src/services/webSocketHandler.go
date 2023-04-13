package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func WebSocketHandler(c *gin.Context) {
    // Upgrade the HTTP connection to a WebSocket connection
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        http.Error(c.Writer, "Failed to upgrade to WebSocket", http.StatusBadRequest)
        return
    }

    // Loop to read messages from the WebSocket connection
    for {
        _, message, err := conn.ReadMessage()
        if err != nil {
            break
        }
        // Handle the incoming message
        fmt.Printf("Received message: %s\n", message)
    }
}
