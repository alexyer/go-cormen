package datastructs

// TODO(alexyer): Implement check for full queue. I'm lazy today :see_no_evil:

type queue struct {
	Data []int
	size int
	head int
	tail int
}

func NewQueue(size int) *queue {
	queue := new(queue)

	queue.Data = make([]int, size)
	queue.size = size
	queue.head = 0
	queue.tail = 0

	return queue
}

func (q *queue) Enqueue(val int) {
	q.Data[q.tail] = val

	if q.tail == q.size-1 {
		q.tail = 0
	} else {
		q.tail++
	}
}

func (q *queue) Dequeue() int {
	val := q.Data[q.head]

	if q.head == q.size-1 {
		q.head = 0
	} else {
		q.head++
	}

	return val
}
