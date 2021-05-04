package main

import (
	"fmt"
)

func f1() int {
	x := 5
	defer func () {
		x++
	}()
	//tmp := x //x是int类型 值传递
	//return tmp
	return x
}

func f2() *int {
	a := 5
	b := &a
	defer func () {
		*b++
	}()
	//temp_data := b
	//return temp_data
	return b
}

func main()  {
	/*
	defer 语句是go体用的一种用于注册延迟调用的机制， 它可以让当前函数执行完毕之后执行
	对于python的with语句来说，
	 */
	//f, _ := os.Open("xxx.txt")
	//defer f.Close()

	fmt.Println("test1")
	//defer之后只能是函数调用 不能是表达式 比如 a++
	defer fmt.Println("defer test1") // 压入栈里面
	defer fmt.Println("defer test2")
	defer fmt.Println("defer test3")
	fmt.Println("test2")
	// 打印输出：
	//test1
	//test2
	//defer test3
	//defer test2
	//defer test1
	//1. 如果有多个defer会出现什么情况 多个defer是按照先入后出的顺序执行

	fmt.Println("===== defer语句执行时的拷贝机制 =====")
	//defer语句执行时的拷贝机制
	test := func () {
		fmt.Println("test1")
	}
	defer test()
	test = func () {
		fmt.Println("test2")
	}
	fmt.Println("test3")
	// test3
	//test1

	//defer语句执行时的拷贝机制
	x := 10
	defer func (a *int) {
		fmt.Println(*a)
	}(&x) // 11
	defer func (a int) {
		fmt.Println(a)
	}(x) // 10
	// 此处的defer函数并没有参数，函数内部使用的值是全局的值
	// 闭包
	defer func () {
		fmt.Println(x)
	}() // 11
	x++

	fmt.Println(f1()) //是不是就意味着 defer中影响不到外部的值呢  // 5
	fmt.Println(*f2()) // 6
	//defer本质上是注册了一个延迟函数，defer函数的执行顺序已经确定
	//defer 没有嵌套 defer的机制是要取代try except finally
	//https://www.cnblogs.com/zhangboyu/p/7911190.html
	//https://studygolang.com/articles/24044?fr=sidebar
}
