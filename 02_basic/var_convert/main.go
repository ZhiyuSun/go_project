package main

import (
	"fmt"
	"strconv"
)

func main()  {
	// 1. 基本的类型转换
	a := int(3.0)
	fmt.Println(a)
	// 在go语言中不支持变量间的隐式类型转换
	// 1. 变量间类型转换不支持
	var b int = 5.0 //常量,常量到变量是会进行隐式转换
	fmt.Println(b)
	c := 5.1
	fmt.Printf("%T\n", c) // float64
	var d int = int(c)
	fmt.Println(d)

	//var e int = int("1") // 这是不行的

	// Go允许在底层结构相同的两个类型之间互转
	var a2 int64 = 56
	var b2 int32 = int32(a2)
	fmt.Println(b2)

	// 不是所有数据类型都能转换的，例如字母格式的string类型"abcd"转换为int肯定会失败
	// 低精度转换为高精度时是安全的，高精度的值转换为低精度时会丢失精度。例如int32转换为int16，float32转换为int
	// 这种简单的转换方式不能对int(float)和string进行互转，要跨大类型转换，可以使用strconv包提供的函数

	// int转字符串 itoa
	fmt.Printf("%T\n", strconv.Itoa(int(64)))
	// 字符串转int atoi
	data, _ := strconv.Atoi("12")
	fmt.Println(data)
	//parse类的函数
	r, err := strconv.ParseBool("False") // 大小写都能转换
	fmt.Println(err)
	fmt.Println(r)
	r2, err2 := strconv.ParseBool("q")
	fmt.Println(err2)
	fmt.Println(r2)
	f, err3 := strconv.ParseFloat("3.1415", 32)
	fmt.Println(err3)
	fmt.Printf("%T\n", f)

	i, err := strconv.ParseInt("012", 0, 64)
	fmt.Println(err)
	fmt.Printf("%T %d\n", i, i)
	u, err := strconv.ParseUint("42", 10, 64)
	fmt.Printf("%T %d\n", u, u)

	//其他类型转字符串
	//s := strconv.FormatBool(true)
	//fmt.Println(s)
	//s := strconv.FormatFloat(3.1415, 'E', -1, 64)
	//fmt.Println(s)
	//s := strconv.FormatInt(-42, 16) //表示将-42转换为16进制数，转换的结果为-2a。
	//s := strconv.FormatUint(42, 10)
	//fmt.Println(s)
	//var data int
	//var err error
	//if data, err = strconv.Atoi("12");err != nil {
	//	fmt.Println("转换出错")
	//}
	//fmt.Println(data)

}
