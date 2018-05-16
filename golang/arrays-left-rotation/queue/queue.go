package queue

type Queue struct {
	Items chan int
}

func New(size int) *Queue {
	if size <= 0{
		return nil
	}
	q := new(Queue)
	q.Items = make(chan int,size)
	return q
}

func (queue *Queue) Enqueue(item int){
	queue.Items <- item
}

func (queue *Queue) Dequeue() int {
	return <- queue.Items
}
