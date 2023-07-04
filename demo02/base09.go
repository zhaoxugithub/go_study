package main

import (
	"fmt"
	"time"
)

/**
 *ClassName base09
 *Description TODO
 *Author 11931
 *Date 2023-07-05:1:05
 *Version 1.0
 **/

func producer1(ch chan string) {
	fmt.Println("producer start")
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
	fmt.Println("producer end")
}

func customer1(ch chan string) {
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}

func main() {
	fmt.Println("main start")
	ch := make(chan string, 1)
	go producer1(ch)
	go customer1(ch)

	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}
