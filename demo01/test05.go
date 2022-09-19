package main

import "fmt"

//这种方式是值传递
func modify1(a [5]int) {
	a[0] = 10
}

//这种方式是引用传递
func modify22(a *[5]int) {
	a[0] = 30
}

func moddify222(a []int) {

}

func main() {

	//申明的是切片
	arrr := []int{2, 3, 4, 5, 6}
	moddify222(arrr)

	//申明的数组
	arrr1 := [...]int{3, 4, 5, 7, 8}

	//数组传参的实参和形参个数要保持一致
	modify1(arrr1)
	fmt.Println(arrr1)

	modify22(&arrr1)
	fmt.Println(arrr1)
}
