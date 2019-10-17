package multicpu

import (
	"sync"
	"sync/atomic"
)

func Count() int64 {

	var count int64
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&count, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	return atomic.LoadInt64(&count)
}
