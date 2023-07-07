package main

import "fmt"

/**
 *ClassName go_defer
 *Description TODO
 *Author 11931
 *Date 2023-07-06:0:13
 *Version 1.0
 **/

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func calcTest() {
	x := 1
	y := 2

	defer calc("A", x, calc("B", x, y))

	x = 3
	defer calc("C", x, calc("D", x, y))
	y = 4
}

// defer 函数定义的顺序 与 实际执的行顺序是相反的，也就是最先声明的最后才执行。
func deferTest() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	fmt.Println("main")
}

func test01() {
	var a = 1
	var b = 2

	/*
		这种是值赋值,defer里面的a,b是按照顺序提前被定义赋值为1,2,只不过没有执行
	*/
	defer fmt.Println(a + b)
	a = 2
	fmt.Println("main")
}

func test02() {
	var a = 1
	var b = 2
	// 这个闭包,闭包获取变量相当于引用传递，而非值传递,所以这个a = 2
	// a 和 b 的值在函数运行时，才能确定。
	defer func() {
		fmt.Println(a + b)
	}()
	a = 2
	fmt.Println("main")
}

func test03() {
	var a = 1
	var b = 2

	// 这种闭包,如果函数体内有参数,属于值传递
	defer func(a int, b int) {
		fmt.Println(a + b)
	}(a, b)
}

// return
func t1() int {
	a := 10
	defer func() {
		a++
	}()
	a = 2
	return a
}

func t2() (a int) {
	defer func() {
		a++
	}()
	return 1
}

func main() {
	test01()
	println("-------------")
	test02()
	println("--------test return-----")

	println(t1())
	println("--------")

	println(t2())
}
