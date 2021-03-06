package multicpu_invalid_test

import (
	"fmt"
	multicpu_invalid "github.com/yutachaos/go-practice/atomic/multicpu"
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
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("GOMAXPROCS: %d", test.GOMAXPROCS), func(t *testing.T) {
			runtime.GOMAXPROCS(test.GOMAXPROCS)
			actual := multicpu_invalid.Count()
			if actual != test.expect {
				t.Errorf("Assert error failed actual: %d expext: %d GOMAXPROCS: %d", actual, test.expect, runtime.GOMAXPROCS(0))
			}
		})
	}

}
