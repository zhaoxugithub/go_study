package main

import "fmt"

//切片是一种动态数组，比数组操作灵活，长度不是固定的，可以进行追加和删除。
//len()和cap()返回的结果可以相同或者不同
func sliceTest() {
	//切片的申明
	var sli_1 []int
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_1), cap(sli_1), sli_1)
	var sli_2 = []int{}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_2), cap(sli_2), sli_2)
	var sli_3 = []int{3, 4, 5, 6, 7}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_3), cap(sli_3), sli_3)
	sli_3 = append(append(append(sli_3, 6), 7), 9)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_3), cap(sli_3), sli_3)
	var sli_4 = make([]int, 5, 5)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_4), cap(sli_4), sli_4)
	sli_5 := make([]int, 9, 10)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_5), cap(sli_5), sli_5)
}

//切片的截取
func sliceSub() {
	sli_1 := []int{}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_1), cap(sli_1), sli_1)
	for i := 0; i < 10; i++ {
		sli_1 = append(sli_1, i)
	}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_1), cap(sli_1), sli_1)
	fmt.Println("sli[1] == ", sli_1[1])
	fmt.Println("sli[:] ==", sli_1[:])
	fmt.Println("sli[1:] ==", sli_1[1:])
	fmt.Println("sli[:4] ==", sli_1[:4])
	//下标是0-2 不包括3
	fmt.Println("sli[0:3] ==", sli_1[0:3])
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_1[0:3]), cap(sli_1[0:3]), sli_1[0:3])

	fmt.Println("sli[0:3:4] ==", sli_1[0:3:4])
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_1[0:3:4]), cap(sli_1[0:3:4]), sli_1[0:3:4])
}

//删除切片
func sliceDelete() {
	sli := []int{}
	for i := 0; i < 10; i++ {
		sli = append(sli, i)
	}
	fmt.Println(sli)
	//删除尾部 2 个元素
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli[:len(sli)-2]), cap(sli[:len(sli)-2]), sli[:len(sli)-2])
	//删除开头 2 个元素
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli[2:]), cap(sli[2:]), sli[2:])
	//删除中间 2 个元素
	sli = append(sli[:3], sli[3+2:]...)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)
}
func main() {
	sliceTest()
	//sliceSub()
	//sliceDelete()
}
