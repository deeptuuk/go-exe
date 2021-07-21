package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"test", 2000, 10, []string{"xingye", "zhangbozhi"}}

	//结构体 -> json
	jsonStr, err := json.Marshal(movie)

	if err != nil {
		fmt.Println("json marshal error")
		return
	} else {
		fmt.Printf("jsonStr = %s\n", jsonStr)
	}

	//json -> 结构体
	//jsonStr = {"title":"test","year":2000,"rmb":10,"actors":["xingye","zhangbozhi"]}

	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error")
		return
	} else {
		fmt.Printf("myMovie = %v\n", myMovie)
	}

}
