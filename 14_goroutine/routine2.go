package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			//return
			runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("main Goroutine exit")
}
