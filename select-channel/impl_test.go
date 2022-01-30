package selectchannel

import (
	"fmt"
	"testing"
	"time"
)

/*
============================select channel=======================-=================================
terkadang terdapat case dimana kita membuat beberapa channel, dan menjalankan beberapa go routines.
lalu kita ingin mendaptkan data dari semua channel tersebut.
untuk melakukan hal tersebut, kita bisa menggunakan select channel digolang.
dengan select channel, kita bisa memilih data tercepat dari beberapa channel,
jika data datang secara bersamaan dibeberapa channel, maka akan dipilih secara random
=================================================================================================== */

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "andi ahmad saputar"
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	// NOTE: variable counter untuk menghindari error deadlock!
	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}

}
