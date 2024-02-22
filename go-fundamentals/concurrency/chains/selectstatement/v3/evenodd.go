package main

import "fmt"

type evenChan chan int
type oddChan chan int
type quitChan chan int

func main() {
	// channels
	even := make(evenChan)
	odd := make(oddChan)
	quit := make(quitChan)

	// a bunch of numbers
	numbers := loadNumbers(20)

	// create a goroutine to send signals to channels
	go senderEvenOddNumbers(even, odd, quit, numbers)

	// receive signals from channels
	receiver(even, odd, quit)
}

func loadNumbers(amount int) []int {
	var numbers []int
	for n := range amount {
		numbers = append(numbers, n)
	}
	return numbers
}

func senderEvenOddNumbers(evenCh evenChan, oddCh oddChan, quitCh quitChan, numbers []int) {
	for n := range numbers {
		if n%2 == 0 {
			evenCh <- n
		} else {
			oddCh <- n
		}
	}
	quitCh <- 66
}

func receiver(evenCh evenChan, oddCh oddChan, quitCh quitChan) {
	var evenNumbers []int
	var oddNumbers []int
EXIT:
	for {
		select {
		case even := <-evenCh:
			evenNumbers = append(evenNumbers, even)
			fmt.Println("Even:\t", even)
		case odd := <-oddCh:
			oddNumbers = append(oddNumbers, odd)
			fmt.Println("Odd:\t", odd)
		case quit, ok := <-quitCh: // comma ok, like in map to verify if a key exists
			fmt.Println("Quit:\t", quit, ok)
			break EXIT // break with label or return, that is the same
		}
	}
	fmt.Println("Even numbers:\t", evenNumbers)
	fmt.Println("Odd numbers:\t", oddNumbers)
}
