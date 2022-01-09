package main

import (
	"fmt"
	"strconv"
	"time"
)

func goFunc() {

	fmt.Println("main start...")
	go func() {
		fmt.Println("thread start ...")
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("main end...")
}

// 申明chan管道
func chanTest() {
	//申明不带缓冲管道
	ch1 := make(chan string)

	//申明带10个缓冲的管道
	ch2 := make(chan string, 10)

	//申明只读管道
	ch3 := make(<-chan string)

	//申明只写管道
	ch4 := make(chan<- string)

	fmt.Println(ch1)
	fmt.Println(ch2)
	fmt.Println(ch3)
	fmt.Println(ch4)
}

// 如果管道不带缓冲，赋值之后就会阻塞
//不带缓冲的通道，进和出都会阻塞。
func chanExeNoBuffer() {

	fmt.Println("main start..")
	//申明一个无缓冲的管道
	ch := make(chan string)

	ch <- "a"
	fmt.Println("管道之后就会被阻。。。。。")
	go func() {
		val := <-ch
		fmt.Println(val)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main end..")
}

// 带缓冲的通道，进一次长度 +1，出一次长度 -1，如果长度等于缓冲长度时，再进就会阻塞。
func chanExeBuffer() {

	fmt.Println("main start..")
	//申明一个无缓冲的管道
	ch := make(chan string, 1)

	ch <- "a"
	fmt.Println("管道之后就会被阻。。。。。")
	go func() {
		val := <-ch
		fmt.Println(val)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main end..")
}

func producer(ch chan string) {
	fmt.Println("producer start ....")
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
		ch <- strconv.Itoa(i)
	}
	fmt.Println("producer end ....")
}

func customer(ch chan string) {
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}

func main() {
	fmt.Println("main start ...")
	ch := make(chan string, 3)

	go producer(ch)
	go customer(ch)

	time.Sleep(1 * time.Minute)

	fmt.Println("main end...")
}
