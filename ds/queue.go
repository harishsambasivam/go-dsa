package ds

type Queue struct {
	values []int
	length int
}

func (queue *Queue) enqueue(value int) int {
	q := *queue
	q.values = append(q.values, 0)
	copy(q.values[1:], q.values)
	q.values[0] = value
	return len(q.values)
}

func (queue *Queue) dequeue() int {
	length := len((*queue).values)
	value := (*queue).values[length-1]
	// *queue = append(*queue[:length-2],[]int{}...)
	return value
}

func main() {
	var queue Queue
	queue.enqueue(1)
	queue.dequeue()
}
