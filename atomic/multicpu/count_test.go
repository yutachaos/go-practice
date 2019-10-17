package multicpu_test

import (
	"fmt"
	"github.com/yutachaos/go-practice/atomic/multicpu"
	"runtime"
	"testing"
)

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
		{
			GOMAXPROCS: 5,
			expect:     1000,
		},
		{
			GOMAXPROCS: 6,
			expect:     1000,
		},
		{
			GOMAXPROCS: 7,
			expect:     1000,
		},
		{
			GOMAXPROCS: 8,
			expect:     1000,
		},
	}
	fmt.Printf("your CPUs: %d", runtime.NumCPU())
	for _, test := range tests {
		t.Run(fmt.Sprintf("GOMAXPROCS: %d", test.GOMAXPROCS), func(t *testing.T) {
			runtime.GOMAXPROCS(test.GOMAXPROCS)
			actual := multicpu.Count()
			if actual != test.expect {
				t.Errorf("Assert error failed actual: %d expext: %d GOMAXPROCS: %d", actual, test.expect, runtime.GOMAXPROCS(0))
			}
		})
	}

}
