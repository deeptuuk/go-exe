package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}

	s1 := s[0:2]

	fmt.Printf("s1 = %v\n", s1)

	s2 := make([]int, 3)

	copy(s2, s)

	s[0] = 100

	fmt.Printf("s = %v\n", s)
	fmt.Printf("s2 = %v\n", s2)

	var test []int
	test = append(test, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(test), cap(test), test)

}
