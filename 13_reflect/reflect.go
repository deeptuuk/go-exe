package main

import "fmt"

func main() {
	var a string
	//pair<statictype:string, value:"aceld">
	a = "aceld"

	//pair<type:string, value:"aceld">
	var allType interface{}
	allType = a

	str, ok := allType.(string)
	if ok {
		fmt.Println("str = ", str)
	}
}
