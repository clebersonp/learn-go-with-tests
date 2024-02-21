package main

import "fmt"

const amountInt = 500

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	var numbers []int
	for i := range amountInt {
		numbers = append(numbers, i)
	}

	firsElements := numbers[:(len(numbers) / 2)]
	go func(nums []int) {
		for _, n := range nums {
			c1 <- n
		}
	}(firsElements)

	lastElements := numbers[(len(numbers) / 2):]
	go func(nums []int) {
		for _, n := range nums {
			c2 <- n
		}
	}(lastElements)

	for range numbers {
		// select case one or case other, case both at the same time one randomly will be chosen
		select {
		case v := <-c1:
			fmt.Println("Received number from c1:\t", v)
		case v := <-c2:
			fmt.Println("Received number from c2:\t", v)
		}
	}
}
