package main

import (
	"fmt"
)

/**
 *ClassName OperatorArray
 *Description TODO
 *Author 11931
 *Date 2023-07-08:15:38
 *Version 1.0
 **/

func initCustomArray1() [3]int {
	array := [...]int{1, 2, 3}
	return array
}

// 这种是值传递
func updateCustomArray1(array [3]int) {
	array[0] = 100
	array[1] = 200
	array[2] = 300
}

func updateCustomArray2(array *[3]int) {
	array[0] = 100
	array[1] = 200
	array[2] = 300
}

// 数组长度初始化就被注定了
func createCustomArray1() {

}

// 删除第i个元素
func deleteCustomArray1(array [3]int, i int) []int {
	return append(array[:i], array[i+1:]...)
}

func deleteCustomArray2(array *[3]int, i int) {
	//append(array[:i], array[i+1:]...) // 报错
}

func main() {
	array := initCustomArray1()
	fmt.Println(array)

	// 这种更新是值传递
	updateCustomArray1(array)
	fmt.Println(array)

	// 通过指针的方式
	updateCustomArray2(&array)
	fmt.Println(array)

	deleteArray := deleteCustomArray1(array, 0)
	fmt.Println(deleteArray)
}
