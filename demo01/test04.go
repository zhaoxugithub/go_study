package main

import "fmt"

func main() {
	//一维数组
	var arra [5]int
	fmt.Println(arra)
	//var ar1 =  [5] int {1, 2}
	var arr_2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr_2)
	arr_3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr_3)
	arr_4 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr_4)
	arr_5 := [5]int{0: 3}
	fmt.Println(arr_5)
	// 二维数组
	bbb := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(bbb)
}
