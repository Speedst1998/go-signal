package websocket

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type ConnectionQueueHandler struct {
	q               Queue
	connection      *websocket.Conn
	mutex           *sync.Mutex
	mediaServerName string
}

func MakeConnectionQueueHandler(connection *websocket.Conn, mediaServerName string) ConnectionQueueHandler {
	list := make([]DescriptionRequest, 0)
	return ConnectionQueueHandler{Queue{list}, connection, &sync.Mutex{}, mediaServerName}
}

func (this ConnectionQueueHandler) Enqueue(channel chan Result, desc string) {
	this.mutex.Lock()
	this.q.enqueue(DescriptionRequest{channel, desc})
	println(this.q.size())
	println(this.q.list)
	if this.q.size() == 1 {
		println("Consuming")
		go this.consume()
	}
	this.mutex.Unlock()
}

func (this ConnectionQueueHandler) consume() {
	println("inside consume")
	for this.q.isNotEmpty() {
		this.mutex.Lock()
		descriptionRequest := this.q.dequeue()
		channel := descriptionRequest.ResultChannel
		description := descriptionRequest.Description
		println("Reading from channel")
		answer, err := exchangeDescription(description, this.connection)
		println("Returning result to channel")
		channel <- Result{Answer: answer, Err: err}
		this.mutex.Unlock()
	}
}

func exchangeDescription(description string, connection *websocket.Conn) (string, error) {
	println("Exchanging Description")

	err := connection.WriteMessage(websocket.TextMessage, []byte(description))
	if err != nil {
		println("WHATIS THE ERROR")
		println(err)
	}
	_, message, err := connection.ReadMessage()
	fmt.Printf("Received message: %s\n", message)
	return string(message), err

}
