package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func count() int64 {

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

func TestCount(t *testing.T) {
	tests := []struct {
		GOMAXPROCS int
		expect     int64
	}{
		{
			GOMAXPROCS: 1,
			expect:     1000,
		},
		{
			GOMAXPROCS: 2,
			expect:     1000,
		},
		{
			GOMAXPROCS: 3,
			expect:     1000,
		},
		{
			GOMAXPROCS: 4,
			expect:     1000,
		},
	}
	fmt.Printf("Your CPUs: %d \n", runtime.NumCPU())
	for _, test := range tests {
		t.Run(fmt.Sprintf("GOMAXPROCS: %d", test.GOMAXPROCS), func(t *testing.T) {
			runtime.GOMAXPROCS(test.GOMAXPROCS)
			actual := count()
			if actual != test.expect {
				t.Errorf("Assert error failed actual: %d expext: %d GOMAXPROCS: %d", actual, test.expect, runtime.GOMAXPROCS(0))
			}
		})
	}

}
