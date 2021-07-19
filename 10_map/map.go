package main

import "fmt"

func main() {
	var myMap1 map[string]string

	if myMap1 == nil {
		fmt.Println("my Map1 equal nil!")
	} else {
		fmt.Println("my Map1 is not equal nil!")
	}

	myMap1 = make(map[string]string, 10)
	myMap1["one"] = "java"
	myMap1["two"] = "c++"
	myMap1["three"] = "python"
	fmt.Println(myMap1)

	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python"
	fmt.Println(myMap2)

	myMap3 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(myMap3)

	//var test map[int]string
	//test = make(map[int]string)
	test := make(map[int]string)
	test[0] = "hello"
}
