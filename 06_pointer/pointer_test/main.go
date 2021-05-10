package main

import "fmt"

func swap1(a int, b int){
	//用于交换a和b
	c := a
	a = b
	b = c
	fmt.Println(a, b)
}

func swap(a *int, b *int){
	//用于交换a和b
	c := *a
	*a = *b
	*b = c
}
func main()  {
	fmt.Println("===== 什么是指针 =====")
	//什么指针，我们提一个问题
	a := 10
	b := 20
	swap1(a, b) // 为什么交换不成功， 这个函数运行完成以后 我想要把a和b的值变掉
	fmt.Println(a, b) // 没有被交换成功，因为是值传递
	//a, b = b, a // 这可以达到交换目的
	fmt.Printf("%p\n", &a) // 变量有地址， 16进制数， 0xc00000a0b8
	//现在有一种特殊的变量类型，这个变量只能保存地址值
	var ip *int //这个变量里面就只能保存地址类型这种值
	ip = &a
	fmt.Println(ip)

	//如果要修改指针指向的变量的值，用法也比较特殊
	*ip = 30 // 顺着指针把a的值改了
	fmt.Println(a) // 30

	//如何定义指针变量 如果修改指针变量指向的内存中的值。 通过指针去取值的时候不知道应该取多大的连续内存空间
	fmt.Printf("ip所指向的内存空间地址是：%p, 内存中的值是: %d\n",ip, *ip)
	swap(&a, &b) // 30, 20
	fmt.Println(a, b) // 20, 30

	//还不足以说服大家
	//但是go中数组是值传递 数组中有100万个值，值传递会造成大量复制，对于这种一般我们都采用切片来传递
	//在python中list和dict这种传递都是引用传递

	//指针还可以指向数组 指向数组的指针 数组是值类型
	arr := [3]int{1,2,3}
	var ip1 *[3]int = &arr
	fmt.Println(ip1)

	//指针数组
	var ptrs [3]*int //创建能够存放三个指针变量的数组
	//很多时候都是函数参数的时候指明的类型
	//指针的默认值是nil
	fmt.Println(ptrs)
	if ip1 != nil {

	}

	//像python和java这种语言都是极力的屏蔽指针， c/c++ 都提供了指针 指针本身是很强大
	//c和c++中指针的功能很强大 指针的转换 指针的偏移 指针的运算
	//go语言没有屏蔽指针，但是go语言在指针上做了大量的限制，安全性高很多，相比较 c和c++灵活性就降低了
	//指针变量中涉及到两个符号 & 和 *

	//make， new， nil

	// temp test
	sun := [...]int{1,2,3}
	zhi := &sun
	fmt.Println(sun)
	fmt.Println(zhi[1])
}
