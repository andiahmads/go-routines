package channelinout

import (
	"fmt"
	"testing"
	"time"
)

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)

	channel <- "andi ahmad saputra"
}

func OnlyOut(channel <-chan string) {

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
	close(channel)

}
