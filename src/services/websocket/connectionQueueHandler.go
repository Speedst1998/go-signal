package websocket

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type ConnectionQueueHandler struct {
	q           Queue
	mediaServer MediaServer
	mutex *sync.Mutex
}

func makeConnectionQueueHandler(mediaServer MediaServer) ConnectionQueueHandler {
	list := make([]chan Result, 0)
	return ConnectionQueueHandler{Queue{list}, mediaServer, &sync.Mutex{} }
}

func (this ConnectionQueueHandler) Enqueue(channel chan Result) {
	this.mutex.Lock()
	this.q.enqueue(channel)
	if this.q.size() == 1 {
		go this.consume()
	}
	this.mutex.Unlock()
}

func (this ConnectionQueueHandler) consume() {
	for(this.q.isNotEmpty()){
		this.mutex.Lock()
		channel := this.q.dequeue()
		result := <- channel
		description := result.value
		exchangeDescription(description, this.mediaServer)
		this.mutex.Unlock()
	}
}

func exchangeDescription(description string, mediaServer MediaServer) (string, error) {
	println("in clineconnect")

	mediaServer.Connection.WriteMessage(websocket.TextMessage, []byte(description))
	println("WHATIS THE ERROR")

	_, message, err := mediaServer.Connection.ReadMessage()
	fmt.Printf("Received message: %s\n", message)
	return string(message), err
    
}
