// Package queue provides a lock-free queue and two-Lock concurrent queue which use the algorithm proposed by Michael and Scott.
// https://doi.org/10.1145/248052.248106.

package queue

import (
	"sync/atomic"
	"unsafe"
)

// LFQueue is a lock-free unbounded queue.
type LFQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

type node struct {
	value interface{}
	next  unsafe.Pointer
}

// NewLFQueue returns an empty queue.
func NewLFQueue() *LFQueue {
	n := unsafe.Pointer(&node{})
	return &LFQueue{head: n, tail: n}
}

// Enqueue puts the given value v at the tail of the queue.
func (q *LFQueue) Enqueue(v interface{}) {
	n := &node{value: v}
	for {
		tail := load(&q.tail)
		next := load(&tail.next)
		// if tail and next consistent?
		if tail == load(&q.tail) {
			if next == nil {
				if cas(&tail.next, next, n) {
					// enqueue is done
					cas(&q.tail, tail, n)
					return
				}
			} else { // tail was not pointing to the last node
				// try to swing tail to the next node
				cas(&q.tail, tail, next)
			}
		}
	}
}

// Dequeue remove and returns the value at the head of the queue
// It returns nil if the queue is empty
func (q *LFQueue) Dequeue() interface{} {
	for {
		head := load(&q.head)
		tail := load(&q.tail)
		next := load(&head.next)
		if head == load(&q.head) { // if head, tail and next consistent?
			if head == tail { // is queue empty or tail falling behind
				if next == nil { // is queue empty?
					return nil
				}
				// tail is falling behind, try to advance it
				cas(&q.tail, tail, next)
			} else {
				// read value before CAS otherwise another dequeue
				v := next.value
				if cas(&q.head, head, next) {
					return v // Dequeue is done, return
				}
			}
		}
	}
}

func load(p *unsafe.Pointer) (n *node) {
	return (*node)(atomic.LoadPointer(p))
}

func cas(p *unsafe.Pointer, old, new *node) (ok bool) {
	return atomic.CompareAndSwapPointer(
		p, unsafe.Pointer(old), unsafe.Pointer(new))
}
