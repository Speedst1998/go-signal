package services

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	localWebsocket "example.com/accounting/src/services/websocket"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocketHandler(c *gin.Context, mediaServerSockets map[string]localWebsocket.MediaServer, mediaServerName string) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.Error(c.Writer, "Failed to upgrade to WebSocket", http.StatusBadRequest)
		return
	}
	queueHandler := localWebsocket.MakeConnectionQueueHandler(conn, mediaServerName)
	mediaServerSockets[mediaServerName] = localWebsocket.MediaServer{QueueHandler: queueHandler, Name: mediaServerName}
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

func ClientConnect(c *gin.Context, MediaServerSockets map[string]localWebsocket.MediaServer, mediaServerName string, description string) (string, error) {

	println("Connecting Client")
	mediaServer, ok := MediaServerSockets[mediaServerName]
	println(mediaServer.Name)

	if ok {
		channel := make(chan localWebsocket.Result, 1)
		println(len(channel))
		println(cap(channel))
		println("Enqueuing channel")
		mediaServer.QueueHandler.Enqueue(channel, description)
		println("Sending result to channel")
		println("Waiting for channel response")
		result := <-channel
		println("Got response : ", result.Answer)
		return result.Answer, result.Err
	}
	return "", errors.New("couldnt find media server")
}
