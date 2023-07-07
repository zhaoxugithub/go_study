package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/**
 *ClassName basic.go
 *Description TODO
 *Author 11931
 *Date 2023-07-06:20:52
 *Version 1.0
 **/

// 自定全局变量
var (
	aa = 3
	ss = "kkk"
	bb = true
)

// 定义变量
func variableZeroValue() {
	var a int
	var s string
	fmt.Println("%d %q\n", a, s)
}

// 初始化变量
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	fmt.Println("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}
func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {

	const (
		// 枚举
		cpp = iota
		_
		python
		golang
		javascript
	)
	const (
		b = 1 << (10 * iota)
		kb
		mg
		gb
		tb
		pb
	)
	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, gb, tb, pb)
}

func main() {
	variableTypeDeduction()
	variableShorter()
	euler()
	triangle()
	consts()
	enums()
}
