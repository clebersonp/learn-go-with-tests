package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	channel := make(chan int)
	quit := make(chan int)

	go receiveFromChannelAndQuit(20, channel, quit)
	go sendToChannel(channel, quit)
	wg.Wait()
}

func receiveFromChannelAndQuit(amount int, channel, quit chan int) {
	defer wg.Done()
	for range amount {
		fmt.Println("Received from channel:\t", <-channel) // received signal from cha
		time.Sleep(100 * time.Millisecond)
	}
	quit <- 66 // send signal to quit chan
}

func sendToChannel(channel, quit chan int) {
	defer wg.Done()
	counter := 1
	for {
		fmt.Println("............................")
		select {
		case channel <- counter: // sending signal to chan
			fmt.Println("Sent to channel:\t", counter)
			counter++
		case v := <-quit: // case quit receives some signal then return this function
			fmt.Println("Quited Signal:\t\t", v)
			return
		}
	}
}
