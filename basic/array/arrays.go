package main

import "fmt"

/**
 *ClassName arrays
 *Description TODO
 *Author 11931
 *Date 2023-07-07:17:49
 *Version 1.0
 **/

// golang里面数组是值类型,而非引用类型
func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	// 定义数组并且设置默认值
	var arr1 [5]int

	// := 这个是申明并且定义数数组
	arr2 := [3]int{1, 2, 3}

	// [...] 这个也是申明定义数组,但是没有定义数组的长度
	arr3 := [...]int{2, 3, 4, 5, 6}

	var grid [3][4]int

	fmt.Println("array definitions:")

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	fmt.Println("printArray(arr1)")

	printArray(arr1)
	fmt.Println("printArray(arr3)")

	printArray(arr3)

	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)
}
