package main

import (
	"fmt"
	"time"
)

func main() {
	go func(a int, b int) bool {
		fmt.Println("a = ", a, ", b = ", b)
		return true
	}(10, 10)

	time.Sleep(5 * time.Second)
	fmt.Println("main Goroutine exit")
}
