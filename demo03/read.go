package main

import (
	"fmt"
	"io"
	"os"
)

/**
 *ClassName read
 *Description TODO
 *Author 11931
 *Date 2023-01-01:20:44
 *Version 1.0
 **/

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	//创建一个长度为num的字节数组
	p := make([]byte, num)
	//调用reader方法读取
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func ReaderExample() {

	/*FOREND:
	//死循环
	for {
		readerMenu()
		var ch string
		//获取键盘输入
		fmt.Scanln(&ch)
		var (
			data []byte
			err  error
		)

		switch strings.ToLower(ch) {
		case "1":
			fmt.Printf("请输入不多于9个字符，以回车结束:")
			data, err = ReadFrom(os.Stdin, 11)
		case "2":
			file, err := os.Open(util.GetProjectRoot() + "/src/chapter01/io/01.txt")
		}
	}*/

}

func readFrom() {
	//从标准输入读取
	from, err := ReadFrom(os.Stdin, 11)
	println(from, err)
	//从普通文件中读取，其中file是os.File的实例
	//ReadFrom()
}

func readerMenu() {
	fmt.Println("")
	fmt.Printf("******** 从不同的来源读取数据 *******")
	fmt.Printf("******** 请选择数据源，请输入：*******")
	fmt.Printf("1 表示 标准输入")
	fmt.Printf("2 表示 普通文件")
	fmt.Printf("3 表示 从字符串")
	fmt.Printf("4 表示 从网络")
	fmt.Printf("b 返回上级菜单")
	fmt.Printf("q 退出")
	fmt.Printf("***********************************")
}

func main() {
	readFrom()
}
