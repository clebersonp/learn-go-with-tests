package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check\t1:", ctx.Err())
	fmt.Println("num goroutines\t1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case ch := <-ctx.Done():
				fmt.Printf("%T, %v\n", ch, ch)
				return
			default:
				n++
				time.Sleep(200 * time.Millisecond)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("error check\t2:", ctx.Err())
	fmt.Println("num goroutines\t2:", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	cancel() // a signal will be sent to <-context.Done() in select statement and application will be finished
	fmt.Println("cancelled context")

	time.Sleep(2 * time.Second)
	fmt.Println("error check\t3:", ctx.Err())
	fmt.Println("num goroutines\t3:", runtime.NumGoroutine())
}
