package main

import (
	"fmt"
	"reflect"
)

/**
 *ClassName initSlices
 *Description TODO
 *Author 11931
 *Date 2023-07-08:21:54
 *Version 1.0
 **/

func getDataType(data interface{}) string {
	val := reflect.ValueOf(data)
	return val.Kind().String()
}

func initSlices1() {
	arr1 := make([]int, 10)
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	arr1[3] = 40
	arr1[4] = 50
	//arr1[5] = 60
	fmt.Printf("type=%v,arr=%v\n", getDataType(arr1), arr1)
}

func initSlices2() {
	var arr1 []int
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	arr1[3] = 40
	arr1[4] = 50
	//arr1[5] = 60
	fmt.Printf("type=%v,arr=%v\n", getDataType(arr1), arr1)
}

func initSlices3() {
	arr1 := []int{}
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	arr1[3] = 40
	arr1[4] = 50
	//arr1[5] = 60
	fmt.Printf("type=%v,arr=%v\n", getDataType(arr1), arr1)
}

func initSlices4() {
	arr1 := []int{}
	fmt.Printf("type=%v,arr=%v\n", getDataType(arr1), arr1)

	// 数组
	arr2 := [...]int{}
	fmt.Printf("type=%v,arr=%v\n", getDataType(arr2), arr2)
}
func main() {
	initSlices1()
	initSlices2()
	initSlices3()
	initSlices4()
}
