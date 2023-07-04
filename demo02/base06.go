package main

import (
	"fmt"
)

func forArrTest() {
	person := [3]string{"value1", "value2", "value3"}
	fmt.Printf("len=%d cap=%d array=%v\n", len(person), cap(person), person)
	//这种遍历key-value
	for k, v := range person {
		fmt.Printf("person[%d]=%s\n", k, v)
	}
	//通过下标遍历
	for i := range person {
		fmt.Printf("person[%d] = %v\n", i, person[i])
	}

	//通过下标遍历
	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d] = %v\n", i, person[i])
	}

	//通过空白符遍历
	for _, name := range person {
		fmt.Printf("person = %v\n", name)
	}
}

func forSliceTest() {
	person := []string{"value1", "value2", "value3"}
	fmt.Printf("len=%d cap=%d array=%v\n", len(person), cap(person), person)
	//这种遍历key-value
	for k, v := range person {
		fmt.Printf("person[%d]=%s\n", k, v)
	}
	//通过下标遍历
	for i := range person {
		fmt.Printf("person[%d] = %v\n", i, person[i])
	}
	//通过下标遍历
	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d] = %v\n", i, person[i])
	}
	//通过空白符遍历
	for _, name := range person {
		fmt.Printf("person = %v\n", name)
	}
}

func forMapTest() {
	map1 := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	for k, v := range map1 {
		fmt.Printf("k = %s,v = %s\n", k, v)
	}
	fmt.Println()
	for i := range map1 {
		fmt.Printf("person[%s] = %s\n", i, map1[i])
	}

	fmt.Println()
	//使用空白符,只遍历了name
	for _, name := range map1 {
		fmt.Println("name :", name)
	}
}

func forSwitchTest() {
	i := 3
	fmt.Printf("i i=%d", i)
	switch i {

	case 1:
		fmt.Println("输出 i =", 1)
	case 2:
		fmt.Println("输出 i =", 2)
	case 3:
		fmt.Println("输出 i =", 3)
		fallthrough
	case 4, 5, 6:
		fmt.Println("输出 i =", "4 or 5 or 6")
	default:
		fmt.Println("输出 i =", "xxx")
	}
}

func main() {
	//forMapTest()
	forSwitchTest()
}
