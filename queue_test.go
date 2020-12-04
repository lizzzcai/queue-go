package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	tests := []struct {
		name  string
		queue Queue
		count int
	}{
		{"lock-free queue", NewLFQueue(), 100},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("queue: %s", tt.name)
		t.Run(testname, func(t *testing.T) {
			count := tt.count
			queue := tt.queue

			for i := 0; i < count; i++ {
				queue.Enqueue(i)
			}

			for i := 0; i < count; i++ {
				v := queue.Dequeue()
				if v == nil {
					t.Fatalf("got a nil value")
				}
				if v.(int) != i {
					t.Fatalf("expect %d, got %d", i, v)
				}
			}
		})
	}
}
