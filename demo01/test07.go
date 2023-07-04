package main

import "fmt"

func initSlice() {

	var slice1 = []int{1, 2, 3}
	ints := append(slice1, 3)
	fmt.Println(ints)
	fmt.Println(slice1)

	ss := make([]int, 4)
	fmt.Println(ss)

	//6 表示切片长度，10表示最大的容量
	slice3 := make([]int, 6, 10)
	fmt.Printf("cap = %d,len = %d,ele = %v\n", cap(slice3), len(slice3), slice3)

	slice3 = append(slice3, 3, 4, 5, 6, 7)
	fmt.Printf("cap = %d,len = %d,ele = %v\n", cap(slice3), len(slice3), slice3)

}

var a int = 0

func addEle() {
	a = a + 1
}

func main() {
	for i := 0; i < 1000; i++ {
		//	//创建线程
		go addEle()
	}
	fmt.Println(a)
	//initSlice()
}
