package hellogoroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestRunHelloWorld(t *testing.T) {
	//create subtest
	t.Run("without go routines", func(t *testing.T) {
		RunHelloWorld()
		fmt.Println("Kocooook")

		time.Sleep(1 * time.Second)
	})

	t.Run("with go routines", func(t *testing.T) {
		go RunHelloWorld()
		fmt.Println("Kocooook")

		time.Sleep(1 * time.Second)
	})

}
