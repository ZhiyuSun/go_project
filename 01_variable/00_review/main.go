package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// GOLD

// 在外面定义变量，只能用var开头
// 这个不是全局变量，而是包内部的变量
// := 只能在函数内使用，在这里只能用var关键字
//var aa = 3
//var ss = "kkk"
//var bb = true
var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	// 为了把空串打印出来
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

// 去掉类型
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

// 更简单的写法。最推荐的做法
// var能不用就不用
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	c := 3+4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi ) + 1)
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i * math.Pi) + 1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	//const filename string = "abc.txt"
	//const a, b = 3, 4
	// 注意，go语言的常量不要大写，大写有特殊含义
	const (
		filename = "abc.txt"
		a, b = 3, 4
	)
	var c int
	// 这里a和b自己变成了float
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		java
		python
		_
		golang
	)
	fmt.Println(cpp, java, python, golang)

	// 自增枚举值
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)

}

func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	// 验证欧拉公式
	euler()

	// 强制类型转换
	triangle()

	// 常量
	consts()

	// 枚举
	enums()
}
