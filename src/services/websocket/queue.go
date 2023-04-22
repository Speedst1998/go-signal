package websocket

type Queue struct {
	list []DescriptionRequest
}

type DescriptionRequest struct {
	ResultChannel chan Result
	Description   string
}

type Result struct {
	Answer string
	Err    error
}

func (q *Queue) enqueue(descriptionRequest DescriptionRequest) {
	println(q.list)
	q.list = append(q.list, descriptionRequest)
	println(q.list)
	println(len(q.list))
}

func (q *Queue) dequeue() DescriptionRequest {
	descriptionRequest := q.list[0]
	q.list = q.list[1:]
	return descriptionRequest
}

func (q *Queue) isEmpty() bool {
	return len(q.list) == 0
}

func (q *Queue) isNotEmpty() bool {
	return !q.isEmpty()
}

func (q *Queue) size() int {
	return len(q.list)
}
