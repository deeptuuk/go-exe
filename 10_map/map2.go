package main

import "fmt"

func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key = ", key, ", value = ", value)
	}
}

func changeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	printMap(cityMap)

	cityMap["USA"] = "DC"

	delete(cityMap, "China")

	fmt.Println("----------")

	printMap(cityMap)

	fmt.Println("----------")

	changeValue(cityMap)

	printMap(cityMap)
}
