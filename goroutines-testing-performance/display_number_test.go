package goroutines_testing_peformance

import (
	"testing"
	"time"
)

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(1 * time.Second)
}
