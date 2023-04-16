package websocket

type ConnectionQueueHandler struct {
	q           Queue
	mediaServer MediaServer
}

func (this ConnectionQueueHandler) Enqueue(channel chan string) {
	//TODO : check for race condition here
	this.q.enqueue(channel)
	if this.q.size() == 1 {
		go this.consume()
	}
}

func (this ConnectionQueueHandler) consume() {
	// if(!this.q.isEmpty())
	// channel := this.q.dequeue()
	// description := <-channel
}
