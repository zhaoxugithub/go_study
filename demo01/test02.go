package main

// 静态变量
func constVar() {
	const str1 string = "ffdfsda"
	const str2 = "bbbb"
	const v1, v2, v3 = "1", "2", "3"
	const vv1, vv2, vv3 string = "1", "2", "3"
}

// 普通变量
func simpleVar() {

	var ff1 int32 = 19
	println(ff1)

	var ff2 = 40
	println(ff2)

	var f33 = "fafsdf"
	println(f33)

	var f1, f2, f3 = 1, 2, 3
	println(f2, f1, f3)

	//自己进行类型判断
	sd := 12
	println(sd)

	sd1, sd2, s3 := 1, 2, 3
	println(sd1, sd2, s3)

	ss1, ss2, ss3 := 2, "sssss", 4
	println(ss3, ss2, ss1)
}

func main() {
	simpleVar()
}
