package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}

		close(c)
	}()

	/*
		for {
			if data, ok := <-c; ok {
				fmt.Println(data)
			} else {
				fmt.Println("channel is shutdow")
				break
			}
		}
	*/

	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("main goroutine is ending...")
}
