package main

import "fmt"

/**
 *ClassName maps
 *Description TODO
 *Author 11931
 *Date 2023-07-08:0:06
 *Version 1.0
 **/

func main() {

	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "immooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil

	fmt.Println("m,m2,m3:")
	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map m")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")

	courseName := m["course"]

	fmt.Println(`m["course"] = `, courseName)

	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key `cause` does not exist")
	}

	fmt.Println("Deleting values")

	name, ok := m["name"]
	fmt.Printf("m[%q] before delete: %q,%v\n", "name", name, ok)
}
