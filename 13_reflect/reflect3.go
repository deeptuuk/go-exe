package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("ReadBook()...")
}

func (this *Book) WriteBook() {
	fmt.Println("WriteBook()...")
}

func main() {
	b := &Book{}
	fmt.Printf("type of b = %T\n", b)

	var r Reader
	fmt.Printf("type of b = %T\n", r)
	r = b
	r.ReadBook()

	var w Writer
	fmt.Printf("type of b = %T\n", w)
	w = b
	w.WriteBook()
}
