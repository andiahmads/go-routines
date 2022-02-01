package syncwaitgroup_test

import (
	"fmt"
	"sync"
	"testing"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	fmt.Println("hello wait group")
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}
	group.Wait()
	fmt.Println("wait group complate")
}
