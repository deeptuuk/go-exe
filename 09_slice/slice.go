package main

import (
	"fmt"
)

func main() {

	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	var slice2 []int
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)

	if slice2 == nil {
		fmt.Println("slice is a empty slice")
	} else {
		fmt.Println("slice is not a empty slice")
	}

	slice2 = make([]int, 3)
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)

	slice2[1] = 100
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)

	var slice3 []int = make([]int, 5)
	fmt.Printf("len = %d, slice = %v\n", len(slice3), slice3)

	slice4 := make([]int, 10)
	fmt.Printf("len = %d, slice = %v\n", len(slice4), slice4)

}
