package main

import "fmt"

// 全局变量，即使被定义了但没有使用，程序依然可以运行
// 全局变量甚至可以和局部变量同名
var c = 20
func main()  {
	//变量的作用域
	//局部变量
	//缩进的习惯
	//静态语言写起来代码多， 但是严谨性很好
	//sex := "Female"
	//if sex == "Female" {
	//	outStr := "女"
	//}
	//
	//fmt.Println(outStr)
	fmt.Printf("%c\n", 97)
	//在go中字符和字符串不是一种类型 字符类型是单引号 字符串是双引号
	//fmt.Printf("%c", 97)
	fmt.Printf("%T\n","孙") // string
	fmt.Printf("%T\n",'孙') // int32
	fmt.Printf("%T\n",'s') // int32
	fmt.Printf("%T","ss") // string
	//a := 1
	//fmt.Println(a)
}
