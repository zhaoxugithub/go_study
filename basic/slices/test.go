package main

import "fmt"

/**
 *ClassName test
 *Description TODO
 *Author 11931
 *Date 2023-07-07:19:12
 *Version 1.0
 **/

func updateSlice2(arr []int) {
	arr[0] = 200
}

//func updateSlice1(arr [...]int) {
//	arr[0] = 300
//}

func updateArr(arr [8]int) {
	arr[0] = 500
}
func updateArrPoint(arr *[8]int) {
	arr[0] = 500
}

func updateSlice3(arr *[]int) {
}

func sliceTest() {

	// 定义的这个是数组,长度是固定的
	sliceArr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	for i, v := range sliceArr {
		fmt.Println(i, v)
	}
	// 更新
	sliceArr[0] = 100
	fmt.Println(sliceArr)
	updateArr(sliceArr)
	fmt.Println(sliceArr) // 这行和37行的输出是一样的

	updateArrPoint(&sliceArr)
	fmt.Println(sliceArr) // 输出,update更新成功
	//updateSlice1(sliceArr) 报错

	// 创建
	//sliceArr[8] = 100 报错
	//apSlice := append(sliceArr, 8)

	// 删除
}

func test() {
	//arr := [...]int{0, 3, 4}
	var arr []int
	nums := 99
	for i := -1; i < nums; i++ {
		//arr[i] = i // 这种方法赋值会报错
		arr = append(arr, i) //
	}
	fmt.Println(arr[:])
}

func main() {
	sliceTest()
}
