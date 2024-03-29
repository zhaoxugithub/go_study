package main

import "fmt"

/**
 *ClassName slices
 *Description TODO
 *Author 11931
 *Date 2023-07-07:18:23
 *Version 1.0
 **/

func updateSlice(s []int) {
	s[0] = 100
}

func main() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	// 表示从下标从2开始,包括2,3,4,5 表示从 6-2 = 4个元素,不包括下标为6的元素
	fmt.Println("arr[2:6] = ", arr[2:6])

	// 总计6个元素,不包括下标为6的元素
	fmt.Println("arr[:6] = ", arr[:6])

	s1 := arr[2:]
	fmt.Println("s1=", s1)

	s2 := arr[:]
	fmt.Println(s2)

	fmt.Println("After updateSlices(s1)")
	updateSlice(s1)

	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlices(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2

	fmt.Println("arr = ", arr)

	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))

	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n", s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)

	s5 := append(s4, 12)

	fmt.Println("s3,s4,s5=", s3, s4, s5)

	fmt.Println("arr = ", arr)

	// uncomment to run sliceOps demo.
	// If we see undefined: slicesOps
	// Please try to run slices.go sliceops.go
	fmt.Println("Uncomment to see sliceOps demo")
}
