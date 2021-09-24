package queue

//Queue define las operaciones para una cola normal.
type Queue interface {
	Enqueue(number int)
	Dequeue() (int, error)
	IsEmpty() bool
	Front() int
}
