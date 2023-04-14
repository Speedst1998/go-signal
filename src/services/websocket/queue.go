package websocket

type Queue struct {
	list []chan string
}

func (q Queue) enqueue(channel chan string) {
	q.list = append(q.list, channel)
}

func (q Queue) dequeue() chan string {
	channel := q.list[0]
	q.list = q.list[1:]
	return channel
}

func (q Queue) isEmpty() bool {
	return len(q.list) == 0
}