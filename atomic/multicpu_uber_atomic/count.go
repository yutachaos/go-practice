package multicpu_uber_atomic

import (
	"go.uber.org/atomic"
	"sync"
)

func count() int64 {

	count := atomic.NewInt64(0)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			count.Add(1)
			wg.Done()
		}()
	}
	wg.Wait()
	return count.Load()
}
