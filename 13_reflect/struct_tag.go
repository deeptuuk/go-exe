package main

import (
	"fmt"
	"reflect"
)

type Resume struct {
	Name string `info:"name" doc:"name"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {

	inputType := reflect.TypeOf(str)
	fmt.Println("======================", inputType)

	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", taginfo, " doc: ", tagdoc)
	}
}

func main() {
	//var re Resume
	re := Resume{Name: "zhang3", Sex: "male"}
	findTag(&re)
}
