package main

import "fmt"

func testMap() {

	map1 := make(map[string]string)
	fmt.Println(map1)

	map1["key1"] = "value1"
	map1["key2"] = "value2"
	map1["key3"] = "value3"
	map1["key4"] = "value4"
	fmt.Println(map1)

	delete(map1, "key1")
	fmt.Println(map1)

	map1["key3"] = "fdfsf"
	fmt.Println(map1)

	for key, value := range map1 {
		fmt.Printf("key = %s,value = %s\n", key, value)
	}

}

func main() {
	testMap()
}
