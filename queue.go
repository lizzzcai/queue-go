package queue

// Queue is a FIFO data structure.
// Enqueue puts a value into its tail,
// Dequeue remoes a value from its head.
type Queue interface {
	Enqueue(v interface{})
	Dequeue() interface{}
}
