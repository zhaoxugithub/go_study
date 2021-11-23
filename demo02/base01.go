package main

import "fmt"

//常量的申明方式
func constVar() {

	//常量
	const name string = "tome"
	fmt.Println(name)

	const age = 30
	fmt.Println(age)

	const field1, filed2 string = "value1", "value2"
	fmt.Println(filed2, field1)

	const field3, field4 = "value1", 40
	fmt.Println(field3, field4)

}

// 普通变量的申明方式
func commonVar() {

	//申明单个变量的三种方式
	var age1 int = 30
	fmt.Println(age1)
	var age2 = 40
	fmt.Println(age2)
	age3 := 50
	fmt.Println(age3)

	//申明的多个变量
	var name1, name2, name3 string = "aa", "bb", "vvv"
	fmt.Println(name1, name2, name3)

	var name4, name5 = "bbbb", 77
	fmt.Println(name4, name5)

	name6, name7 := 50, "ddd"
	fmt.Println(name6, name7)

}

// 输出类型
func fmtOut() {
	fmt.Printf("输出到控制台")
	fmt.Println("----------")
	fmt.Println("输出到控制台并且换行")
	fmt.Printf("name = %s,age=%d\n", "tom", 50)
	fmt.Printf("name = %s,age = %d,height = %v\n", "lisi", 40, fmt.Sprintf("%.2f", 189.765))
}

func main() {
	//constVar()
	commonVar()
	fmtOut()
}
