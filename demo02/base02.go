package main

import "fmt"

//go中的数组是还不可以动态变化的，一旦申明了，长度就是固定的,而且是值传递
//同样长度的数组才可以相互赋值
func arrayTest() {
	var arr_1 = [5]int{2, 3, 4, 5, 6}
	fmt.Println(arr_1)
	arr_1[4] = 9
	fmt.Println(arr_1)

	//值不会改变
	modifyValue(arr_1)
	fmt.Println(arr_1)

	//传递引用变量
	modifyValue2(&arr_1)
	fmt.Println(arr_1)

	var arr_2 = [4]int{}
	fmt.Println(arr_2)

	var arr_3 [5]int
	fmt.Println(arr_3)

	arr_4 := [100]int{}
	arr_4[0] = 1
	arr_4[1] = 2
	fmt.Println(arr_4)

}

func modifyValue(a [5]int) {
	a[0] = 9999
}

func modifyValue2(a *[5]int) {
	a[0] = 9999
}

func main() {
	arrayTest()
}
