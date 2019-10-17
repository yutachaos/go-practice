package singlecpu_test

import (
	"fmt"
	"github.com/yutachaos/go-practice/atomic/singlecpu"
	"runtime"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		expect int64
	}{
		{
			expect: 1000,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("expect: %d", test.expect), func(t *testing.T) {
			actual := singlecpu.Count()
			if actual != test.expect {
				t.Errorf("Assert error failed actual: %d expext: %d GOMAXPROCS: %d", actual, test.expect, runtime.GOMAXPROCS(0))
			}
		})
	}

}
