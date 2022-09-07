package main

import (
	"fmt"
	"math"
	"reflect"
)

/**
 *ClassName base01
 *Description go 基本类型
 *Author 11931
 *Date 2022/9/6:22:06
 *Version 1.0
 **/
func main() {
	test6()
}

func test1() {
	a, b, c := 0b1010, 0o144, 0x64
	fmt.Printf("0b%b,%#o,%#x\n", a, b, c)
	//-128,127
	fmt.Println(math.MinInt8, math.MaxInt8)
}

//用下划线作为数据分隔符，以便于阅读
func test2() {
	a := 111_22_3_44
	println(a)
}

/**
nil 是预定义标识符，不是关键字

1.其含义是默认零值，而非空
2.没有类型，不能作为初始值公编译器推断
3.不同类型的nil 值不能比较
*/
func test3() {
	//x := nil
	//var x1 int = nil
	//println(x1)
}

func test4() {
	//var m map[string]int = nil
	//var c []int = nil
	// 不同类型的nil 值不能比较
	//println(m == c)
}

//标准库 sync/atomic 提供了Bool,Int32 / Uint32、Int64/UInt64,Uintptr 原子操作版本
func test5() {
	var x int
	//0
	println(x)
	//var x1 atomic.Int64
	//go func() {
	//	x1.Add(1)
	//}()
	//
	//x1.Store(x1.Load() + 1)
	//select {}
}

//别名
type X = int

func test6() {
	var a int = 100
	var b X = a
	//int,100
	fmt.Printf("%T,%v\n", b, b)
	//int
	fmt.Println(reflect.TypeOf(b))
}
