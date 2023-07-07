package main

import "fmt"

/**
 *ClassName arrays
 *Description TODO
 *Author 11931
 *Date 2023-07-07:17:49
 *Version 1.0
 **/

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
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
