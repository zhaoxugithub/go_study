package common

import (
	"fmt"
	"reflect"
)

/**
 *ClassName JugeDataType
 *Description TODO
 *Author 11931
 *Date 2023-07-08:22:08
 *Version 1.0
 **/

func isArray(data interface{}) bool {
	val := reflect.ValueOf(data)
	kind := val.Kind()
	return kind == reflect.Array
}

func getDataType(data interface{}) string {
	val := reflect.ValueOf(data)
	return val.Kind().String()
}

func main() {
	var arr [3]int
	var slice []int

	fmt.Println(isArray(arr))   // 输出: true
	fmt.Println(isArray(slice)) // 输出: false

	fmt.Println(getDataType(arr))
	fmt.Println(getDataType(slice))
}
