package main

import (
	"fmt"
	"reflect"
)

/**
 *ClassName arrays2
 *Description TODO
 *Author 11931
 *Date 2023-07-08:15:19
 *Version 1.0
 **/

func getDataType(data interface{}) string {
	val := reflect.ValueOf(data)
	return val.Kind().String()
}

func initArray1() {
	// 定义数组
	arr1 := [5]int{}
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	arr1[3] = 40
	arr1[4] = 50
	//arr1[5] = 60 报错
	fmt.Printf("type=%v,arr=%v\n", getDataType(arr1), arr1)
}

func initArray2() {
	// 定义数组
	var arr1 [5]int
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	arr1[3] = 40
	arr1[4] = 50
	//arr1[5] = 60
	fmt.Printf("type=%v,arr=%v\n", getDataType(arr1), arr1)
}

func initArray3() {
	// 定义数组
	arr1 := [...]int{1, 2, 3, 4, 5}
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	arr1[3] = 40
	arr1[4] = 50
	//arr1[5] = 60
	fmt.Printf("type=%v,arr=%v\n", getDataType(arr1), arr1)
}

// 这种是初始化切片
func initArray4() {
	// 定义数组
	arr1 := make([]int, 10)
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	arr1[3] = 40
	arr1[4] = 50
	//arr1[5] = 60
	fmt.Printf("type=%v,arr=%v\n", getDataType(arr1), arr1)
}

func initArray() {
	initArray1()
	initArray2()
	initArray3()
	initArray4()
}

func main() {
	initArray()
}
