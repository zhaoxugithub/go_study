package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

/**
 *ClassName go_fun_test
 *Description TODO
 *Author 11931
 *Date 2023-07-04:20:08
 *Version 1.0
 **/

type Result1 struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func toJson1(res *Result1) {
	marshal, err := json.Marshal(res)
	if err != nil {
		fmt.Println("json marshal error:", err)
	}
	fmt.Println("json data :", string(marshal))
}

func setData1(res *Result1) {
	res.Code = 500
	res.Message = "fail"
}

// MD5方法
func MD5COpy(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

func getTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func getTimeInt() int64 {
	return time.Now().Unix()
}

func createSign(params map[string]interface{}) string {
	var key []string
	var str = ""
	for s := range params {
		key = append(key, s)
	}
	for k, v := range params {
		println(k)
		println(v)
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params[key[i]])
		} else {
			str = str + fmt.Sprintf("&x1_%v=%v", key[i], params[key[i]])
		}
	}
	// 自定义秘钥
	var secret = "123456789"
	// 自定义签名算法
	sign := MD5COpy(str) + MD5COpy(secret)
	return sign
}

func main() {
	var res Result1
	res.Code = 200
	res.Message = "message"
	toJson1(&res)
	setData1(&res)
	toJson1(&res)

	println("-----------")

	str := "12345"
	fmt.Println("MD5(%s):%s ", str, MD5COpy(str))

	println("-----------------")

	println(getTimeStr())

	println("--------------")

	println(getTimeInt())

	println("---------------")
	params := map[string]interface{}{
		"name": "tom",
		"pwd":  "123456",
		"age":  30,
	}

	fmt.Println("sign:%s\n", createSign(params))
}
