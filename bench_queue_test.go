package queue

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync/atomic"
	"testing"
)

func BenchmarkQueue(b *testing.B) {

	queues := []struct {
		name  string
		queue Queue
	}{
		{"lock-free queue", NewLFQueue()},
	}

	length := 1 << 12
	inputs := make([]int, length)
	for i := 0; i < length; i++ {
		inputs = append(inputs, rand.Int())
	}

	for _, cpus := range []int{4, 32, 1024} {
		runtime.GOMAXPROCS(cpus)
		for _, q := range queues {
			b.Run(fmt.Sprintf("%s#%d", q.name, cpus), func(b *testing.B) {
				b.ResetTimer()

				var c int64
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						i := int(atomic.AddInt64(&c, 1)-1) % length
						v := inputs[i]
						if v >= 0 {
							q.queue.Enqueue(v)
						} else {
							q.queue.Dequeue()
						}
					}
				})
			})
		}
	}
}
