package rangechannel

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRangeChannel(t *testing.T) {

	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "iterasi ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("menerima data = ", data)
	}
	fmt.Println("selesai")

}
