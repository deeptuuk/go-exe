package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	//如何区分传入的参数的类型
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", value)
		fmt.Printf("value type is %T\n", value)
	}
}

type Book struct {
	name string
	auth string
}

func main() {
	book := Book{name: "c++", auth: "zhang3"}

	myFunc(book)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)
}
