package main

import (
	"fmt"
)

func func1() {
	fmt.Println("this is func1")
}

func func2() {
	fmt.Println("this is func2")
}

func func3() {
	fmt.Println("this is func3")
}

func defer_call() {
	defer func1()
	defer func2()
	defer func3()
}

func test_call_defer() int {
	fmt.Println("this is test_call_defer...")
	return 200
}

func test_call_return() int {
	fmt.Println("this is test_call_return...")
	return 100
}

func test_defer_and_return() int {
	defer test_call_defer()
	return test_call_return()
}

func main() {
	//压栈方式
	defer fmt.Println("main end 1")
	defer fmt.Println("main end 2")

	fmt.Println("go 1")
	fmt.Println("go 2")

	defer_call()

	temp := test_defer_and_return()
	fmt.Println("temp = ", temp)
}
