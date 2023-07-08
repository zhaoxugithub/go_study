package main

import "fmt"

/**
 *ClassName OperatorSlices
 *Description TODO
 *Author 11931
 *Date 2023-07-08:23:41
 *Version 1.0
 **/

func main() {
	var slice1 []int
	for i := 0; i < 10; i++ {
		slice1 = append(slice1, i)
	}
	//slice1=[0 1 2 3 4 5 6 7 8 9],len=10,cap=16
	fmt.Printf("slice1=%v,len=%d,cap=%d\n", slice1, len(slice1), cap(slice1))

	s2 := slice1[7:9]
	//s2=[7 8],len=2,cap=9
	fmt.Printf("s2=%v,len=%d,cap=%d\n", s2, len(s2), cap(s2))

	s3 := append(s2, 100)
	//s3=[7 8 100],len=3,cap=9
	fmt.Printf("s3=%v,len=%d,cap=%d\n", s3, len(s3), cap(s3))

	s4 := append(s3, 200)
	//s4=[7 8 100 200],len=4,cap=9
	fmt.Printf("s4=%v,len=%d,cap=%d\n", s4, len(s4), cap(s4))

	s2c1 := slice1[:]
	//
	fmt.Printf("s2c1=%v,len=%d,cap=%d\n", s2c1, len(s2c1), cap(s2c1))
}
