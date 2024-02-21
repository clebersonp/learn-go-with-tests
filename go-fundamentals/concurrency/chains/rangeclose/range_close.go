package main

import "fmt"

func main() {
	c := make(chan int) // chan without buffer
	go sendToChannel(c, 20)
	receiveFromChannel(c)
}

// sendToChannel - c chan<- int i.e: It sends int to chan
func sendToChannel(c chan<- int, amount int) {
	for i := 0; i < amount; i++ {
		c <- i
		fmt.Printf("Sent value to channel:\t\tchan <- %d\n", i+1)
	}
	close(c) // it needs to clone to receiver range stop waiting for
}

// receiveFromChannel - c <-chan int i.e: It receives int from chan
func receiveFromChannel(c <-chan int) {
	for v := range c {
		fmt.Printf("Received value from channel:\t%d <-chan\n", v)
	}
}
