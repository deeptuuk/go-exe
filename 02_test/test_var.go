package main

import (
	"fmt"
)

//全局变量
var g_a int = 400
var g_b = 500
var g_c int

//不支持 :=
//test := "hello"

func main() {
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var c = 200
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var d string = "abcd"
	fmt.Println("d = ", d)
	fmt.Printf("type of d = %T\n", d)

	var e = "zxcv"
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	f := 300 //既初始化， 又赋值
	fmt.Println("f = ", f)
	fmt.Printf("type of e = %T\n", f)

	g := 3.14
	fmt.Println("g = ", g)
	fmt.Printf("type of g = %T\n", g)

	fmt.Println("g_a = ", g_a)
	fmt.Printf("type of g_a = %T\n", g_a)

	fmt.Println("g_b = ", g_b)
	fmt.Printf("type of g_b = %T\n", g_b)

	fmt.Println("g_c = ", g_c)
	fmt.Printf("type of g_c = %T\n", g_c)

	var xx, yy int = 1, 2
	fmt.Println("xx = ", xx, ", yy = ", yy)

	var kk, ll = 100, "testt"
	fmt.Println("kk = ", kk, ", ll = ", ll)

	var (
		vv int  = 100
		jj bool = true
	)

	fmt.Println("vv = ", vv, ", jj = ", jj)
	fmt.Printf("type of vv = %T\n", vv)
	fmt.Printf("type of jj = %T\n", jj)
}
