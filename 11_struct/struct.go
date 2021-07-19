package main

import "fmt"

type myint int

type Book struct {
	title string
	auth  string
}

func printBook(book Book) {
	fmt.Println(book)
}

func changeBook(book *Book) {
	book.auth = "test"
	(*book).title = "hhhhhhh"
}

func main() {
	var a myint = 10
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhang3"

	printBook(book1)
	fmt.Printf("type of book1 = %T\n", book1)

	changeBook(&book1)
	printBook(book1)
}
