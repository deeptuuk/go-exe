package main

import (
	"fmt"

	"github.com/deeptuuk/go_modules_test/lib1"
	"github.com/deeptuuk/go_modules_test/lib2"
)

func main() {
	lib1.Lib1Test()
	lib2.Lib2Test()
	fmt.Println("main exit")
}
