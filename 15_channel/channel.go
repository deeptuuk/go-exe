package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine is ending...")
		fmt.Println("goroutine is running...")

		c <- 666
	}()

	num := <-c
	fmt.Println("num = ", num)
	fmt.Println("main Goroutine exit")
}
