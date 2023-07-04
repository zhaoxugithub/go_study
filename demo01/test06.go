package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

// 获取携程ID
func goID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func goFunc(i int) {
	sprintf := fmt.Sprintf("i=%d,threadId=%v", i, goID())
	println(sprintf)
}

func main() {
	for i := 0; i < 10000; i++ {
		go goFunc(i)
		//fmt.Println(i)
	}
	time.Sleep(time.Second)
}
