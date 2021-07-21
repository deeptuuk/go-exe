package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (m User) Call() {
	fmt.Println("user is called...")
	fmt.Printf("%v\n", m)
}

func main() {
	user := User{1, "Aceld", 18}

	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}) {

	inputType := reflect.TypeOf(input)
	fmt.Println("======================", inputType)
	fmt.Println("inputType is :", inputType.Name())

	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	fmt.Println("======================", inputType.NumMethod())

	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
		fmt.Println("======================")
	}
}
