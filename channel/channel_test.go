package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {

	channel := make(chan string)
	channel2 := make(chan int)
	defer close(channel)
	defer close(channel2)

	//buat go routines dengan anonymous function

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Andi ahmad"

		fmt.Println("Selesai mengirim data ke channel")
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- 05

		fmt.Println("Selesai mengirim data ke channel2")
	}()

	//coba ambil data dari channel
	data := <-channel
	data2 := <-channel2
	fmt.Println(data)
	fmt.Println(data2)

	time.Sleep(5 * time.Second)

}
