package main

import (
	"encoding/json"
	"fmt"
)

//Name 开头
type Person struct {
	Name string
	Age  int
}

//申明结构体
func initStruct() {
	var p1 Person
	p1.Name = "zhangsan"
	p1.Age = 30

	fmt.Println("p1=", p1)

	var p2 = Person{
		Name: "赵旭",
		Age:  30,
	}
	fmt.Println("p2=", p2)

	p3 := Person{Name: "李四", Age: 90}
	fmt.Println("p3=", p3)

	p4 := struct {
		Name string
		Age  int
	}{Name: "sali", Age: 40}
	fmt.Println(p4)
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

//
func setData(res *Result) {
	res.Code = 500
	res.Message = "fail"
}

func toJson(res *Result) {

	jsons, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("json Marshal error: %s", err)
	}
	fmt.Printf("this json is %s\n", string(jsons))
}

func testResult() {
	res := Result{Code: 200, Message: "success"}
	toJson(&res)

	setData(&res)
	fmt.Println(res)

	toJson(&res)
}

func test() {
	var res Result
	res.Code = 200
	res.Message = "success"

	//序列化
	//将对象转成json
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	//将json转成str
	fmt.Println("json data :", string(jsons))

	//反序列化
	var res2 Result
	errs = json.Unmarshal(jsons, &res2)
	if errs != nil {
		fmt.Println("json unmarshal error:", errs)
	}
	fmt.Println("res2 :", res2)
}

func main() {
	test()
}
