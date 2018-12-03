package chapter5

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type counterA struct {
	i int64
}

func (c *counterA) increment() {
	atomic.AddInt64(&c.i, 1)
}

func (c *counterA) cnt() int64 {
	return c.i
}

func DoAtomic() int64 {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counterA{i: 0}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.increment()
		}()
	}

	wg.Wait()

	return c.cnt()
}
