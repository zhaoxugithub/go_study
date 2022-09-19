package main

import (
	"fmt"
	"time"
)

func goFunc(i int) {
	fmt.Println(i)
}

func main() {
	for i := 0; i < 10000; i++ {
		go goFunc(i)
		//fmt.Println(i)
	}
	time.Sleep(time.Second)
}
