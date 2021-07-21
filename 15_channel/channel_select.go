package main

import (
	"fmt"
	"time"
)

func fibonacii(c, quit chan int) {
	x, y, temp := 1, 1, 0

	for {
		select {
		case c <- x:
			temp = x
			x = y
			y = temp + x
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 30; i++ {
			fmt.Println(<-c)
			time.Sleep(1 * time.Second)
		}

		quit <- 0
	}()

	fibonacii(c, quit)
}
