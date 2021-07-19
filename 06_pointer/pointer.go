package main

import "fmt"

func changeValue(p *int) {
	*p = 10
}

func changeTwo(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 1
	fmt.Println("a = ", a)
	changeValue(&a)
	fmt.Println("a = ", a)

	one := 4
	two := 5

	fmt.Println("one = ", one, ", two = ", two)

	changeTwo(&one, &two)

	fmt.Println("one = ", one, ", two = ", two)
}
