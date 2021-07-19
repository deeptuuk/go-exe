package main

import (
	_ "05_init/lib1" // _ 是匿名的，无法使用里面的方法，但可以使用init()
	mylib2 "05_init/lib2"
)

func main() {
	//lib1.Lib1Test()
	//lib2.Lib2Test()
	mylib2.Lib2Test()
}
