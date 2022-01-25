package bufferdchannel

import (
	"fmt"
	"testing"
	"time"
)

func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {

		channel <- "ANDI"
		channel <- "AHMAD"
		channel <- "SAPUTRA"
	}()

	go func() {

		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)

	}()
	time.Sleep(2 * time.Second)
	fmt.Println("selesai")
}
