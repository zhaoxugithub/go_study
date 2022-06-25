package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

//go中的函数申明
func functionName(input1 int, input2 int) (return1 int, return2 int) {
	return return1, return2
}

//传递参数时，将参数复制一份传递到函数中，对参数进行调整后，不影响参数值。
//
//Go 语言默认是值传递。
func funTest() {
}
func MD5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
func getTimestr() string {
	return time.Now().Format("YYYY-MM-dd HH:mm:ss")
}
func getTimeUnix() int64 {
	return time.Now().Unix()
}
func main() {
	str := "12345"
	fmt.Printf("str %s md5 is %s\n", str, MD5(str))
	fmt.Println(getTimestr())
	fmt.Println(getTimeUnix())
}
