package websocket

type Queue struct {
	list []chan Result
}

type Result struct {
	Value string
	Err   error
}

func (q Queue) enqueue(channel chan Result) {
	println(q.list)
	q.list = append(q.list, channel)
	println(q.list)
	println(len(q.list))
}

func (q Queue) dequeue() chan Result {
	channel := q.list[0]
	q.list = q.list[1:]
	return channel
}

func (q Queue) isEmpty() bool {
	return len(q.list) == 0
}

func (q Queue) isNotEmpty() bool {
	return !q.isEmpty()
}

func (q Queue) size() int {
	return len(q.list)
}
