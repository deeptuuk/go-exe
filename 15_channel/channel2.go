package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)
	//fmt.Println("len(c) = ", len(c), ", cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("sub goroutine is ending...")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("sub goroutine is running len(c) = ", len(c), ", cap(c) = ", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("main goroutine receive ", num)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("main goroutine is ending...")
}
