package main

import "fmt"

/**
 *ClassName test
 *Description TODO
 *Author 11931
 *Date 2023-07-07:19:12
 *Version 1.0
 **/

func main() {
	//arr := [...]int{1, 3, 4}
	var arr []int
	nums := 100
	for i := 0; i < nums; i++ {
		//arr[i] = i // 这种方法赋值会报错
		arr = append(arr, i) //
	}
	fmt.Println(arr[:])
}
