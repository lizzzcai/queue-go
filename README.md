# queue-go
Implementation of lock-free queue

# Queue
In computer science, a queue is a collection of entities that are maintained in a sequence and can be modified by the addition of entities at one end of the sequence and the removal of entities from the other end of the sequence. By convention, the end of the sequence at which elements are added is called the back, tail, or rear of the queue, and the end at which elements are removed is called the head or front of the queue, analogously to the words used when people line up to wait for goods or services.

The package defines the following interface:
```go
type Queue interface {
	Enqueue(v interface{})
	Dequeue() interface{}
}
```

# Implementation
* lock-free queue: `LFQueue`

# Run testing
```bash
go test -v
```

# Run benchmark
```bash
go test -bench=.
```

# Reference
* [Michael, M.M. and Scott, M.L., 1996, May. Simple, fast, and practical non-blocking and blocking concurrent queue algorithms. In Proceedings of the fifteenth annual ACM symposium on Principles of distributed computing (pp. 267-275).](https://www.cs.rochester.edu/u/scott/papers/1996_PODC_queues.pdf)
