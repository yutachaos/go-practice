package multicpu_invalid

import (
	"sync"
)

func Count() int64 {

	var count int64
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			count++
			wg.Done()
		}()
	}
	wg.Wait()
	return count
}
