package main

import (
	"fmt"
	"io/ioutil"
)

/**
 *ClassName branch
 *Description TODO
 *Author 11931
 *Date 2023-07-07:1:10
 *Version 1.0
 **/

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"

	}
	return g
}

func main() {
	const fileName = "D:\\document\\go\\demo1\\basic\\branch\\abc.txt"
	if content, err := ioutil.ReadFile(fileName); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("%s\n", content)
	}
}
