package main

import (
	"encoding/json"
	"fmt"
)

//申明Map
func initMap() {

	var map1 map[string]string
	fmt.Println(map1)

	var map2 = map[string]string{}
	fmt.Println(map2)

	var map3 map[string]string = make(map[string]string)
	fmt.Println(map3)

	var map4 = make(map[string]string)
	fmt.Println(map4)

	map5 := map[string]string{"key1": "value1"}
	fmt.Println(map5)

	map6 := make(map[string]string)
	map6["key2"] = "value2"
	fmt.Println(map6)
}

//Map --> Json
func mapToJson() {

	map1 := make(map[string]interface{})
	map1["code"] = 200
	map1["message"] = "success"
	map1["data"] = map[string]interface{}{
		"username": "zhangsan",
		"age":      30,
		"hoppy":    []string{"读书,爬山"},
	}

	//序列化
	jsons, errors := json.Marshal(map1)
	if errors != nil {
		fmt.Println(errors)
	}
	fmt.Println("")
	fmt.Println("--- map to json ---")
	fmt.Println("json data :", string(jsons))

	// 反序列化
	res2 := map[string]interface{}{}
	errors = json.Unmarshal([]byte(jsons), &res2)
	if errors != nil {
		fmt.Println(errors)
	}
	fmt.Println("")
	fmt.Println("--- json to map ---")
	fmt.Println("map data :", res2)

}

func deleteMap() {

	map1 := map[int]interface{}{
		1: "value1",
		2: "value2",
		3: "value3",
	}
	fmt.Println(map1)

	delete(map1, 2)
	fmt.Println(map1)

	map1[2] = "value33333"
	map1[4] = "fdfdsfasfs"
	fmt.Println(map1)

}

func main() {
	initMap()
	//deleteMap()
}
