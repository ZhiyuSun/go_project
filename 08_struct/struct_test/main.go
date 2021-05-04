package main

import (
	"fmt"
	"unsafe"
)

type Course struct {
	Name string
	Price int
	Url string
	// 只能定义数据，聚合数据，不能定义函数
}

//函数的接收者
func (c Course) printCourseInfo()  {
	fmt.Printf("课程名:%s, 课程价格: %d, 课程的地址:%s\n", c.Name, c.Price, c.Url)
}

func (c Course) setPrice1(price int){
	c.Price = price
}
func (c *Course) setPrice(price int){
	c.Price = price
}

//1. 结构体的方法只能和结构体在同一个包中
//2. 内置的int类型不能加方法
func main()  {
	//go语言不支持面向对象
	//面向对象的三个基本特征： 1. 封装 2. 继承 3. 多态。go语言的效果能支持它们
	//4. 方法重载 5. 抽象基类
	// 封装：约束类的行为。go如果要想达到这个封装的效果 那么就一定要解决 1. 变量的封装 2. 方法的封装
	//定义struct go语言没有class这个概念 所以说对于很多人来说会少理解很多面向对象抽象的概念

	//1. 实例化- kv形式
	// 用大括号实例化
	// 后面一定要加逗号
	var c Course = Course{
		Name: "django",
		Price: 100,
		Url: "https://www.baidu.com",
	}
	//访问
	fmt.Println(c.Name, c.Price, c.Url)
	fmt.Println(c)

	//大小写在go语言中的重要性 可见性
	//一个包中的变量或者结构体如果首字母是小写 那么对于另一个包不可见
	//结构体定义的 名称 以及属性首字母大写很重要

	//2. 第二种实例化方式 - 顺序形式
	// 两种模式不能混合使用
	c2 := Course{"scrapy", 110, "https://www.baidu.com"}
	fmt.Println(c2.Name, c2.Price, c2.Url)
	fmt.Printf("%T\n", c2)

	//3. 如果一个指向结构体的指针, 通过结构体指针获取对象的值， 让很多人莫名其妙
	c3 := &Course{"tornado", 100, "https://www.baidu.com"}
	fmt.Printf("%T\n", c3)
	//应该能看出来 go语言实际上在借鉴动态语言的特性 - 很多地方不管如何写都是正确的
	//另一个根本的原因 - go语言的指针是受限的(比如C的指针偏移被屏蔽了)
	fmt.Println((*c3).Name, c3.Price, c3.Url) //这里其实是go语言的一个语法糖 go语言内部会将c3.Name转换成 (*c3).Name

	//4. 零值 如果不给结构体赋值， go语言会默认给每个字段采用默认值
	c4 := Course{}
	fmt.Println(c4.Name, c4.Price, c4.Url)

	//5. 多种方式零值初始结构体
	var c5 Course = Course{}
	fmt.Println(c5.Name, c5.Price, c5.Url)
	var c6 Course
	fmt.Println(c6.Name, c6.Price, c6.Url)
	var c7 *Course = &Course{}
	fmt.Println(c7.Name, c7.Price, c7.Url)
	var c8 *Course = new(Course)
	fmt.Println(c8.Name, c8.Price, c8.Url)
	//var c9 *Course // 会报错
	//fmt.Println(c9.Name)
	//为什么c6和c9表现出来的结果不一样 指针如果只申明不赋值 默认值是nil c6不是指针 是结构体的类型
	//slice map

	//6. 结构体是值类型
	c10 := Course{"scrapy", 110, "https://www.baidu.com"}
	c11 := c10
	fmt.Println(c10)
	fmt.Println(c11)
	c10.Price = 200
	fmt.Println(c10) // 200
	fmt.Println(c11) // 100, 互不影响

	//go语言中struct无处不在
	//7. 结构体的大小 占用内存的大小 可以使用sizeof来查看对象占用的类型
	fmt.Println(unsafe.Sizeof(1)) // int64
	fmt.Printf("%T\n", 1) // int
	//go语言string的本质 其实string是一个结构体
	//type string struct {
	//	Data uintptr // 指针占8个长度
	//	Len int // 长度64位系统占8个长度
	//}
	fmt.Println(unsafe.Sizeof("")) // 8
	fmt.Println(unsafe.Sizeof(true)) // 1
	fmt.Printf("%T\n", true) // bool
	fmt.Println(unsafe.Sizeof(c10)) // 40=16+8+16

	//8. slice的大小
	type slice struct {
		array unsafe.Pointer  // 底层数组的地址
		len int // 长度
		cap int // 容量
	}

	s1 := []string{"django", "tornado", "scrapy", "celery", "snaic", "flask"}
	fmt.Println("切片占用的内存:", unsafe.Sizeof(s1)) // 24


	m1 := map[string]string {
		"s1": "django",
		"s2": "tornado",
		"s3": "scrapy",
		"s4": "celery",
	}
	fmt.Println(unsafe.Sizeof(m1)) // 8

	fmt.Println("===== 结构体方法 =====")
	//结构体方法， 达到了封装数据和封装方法的效果
	c12 := Course{"scrapy", 110, "https://www.baidu.com"}
	Course.printCourseInfo(c12)
	c12.printCourseInfo()
	c12.setPrice1(1000)
	fmt.Println(c12.Price) // 仍然为110，  为什么呢？ 语法糖 函数参数的传递是怎么传递的？ 结构体是值传递
	Course.setPrice1(c12, 1000)
	fmt.Println(c12.Price)
	c12.setPrice(2000)
	fmt.Println(c12.Price)
	(&c12).setPrice(1000) //修改c10的price? 为什么呢？ 语法糖 函数参数的传递是怎么传递的？ 结构体是值传递
	fmt.Println(c12.Price)

	//结构体的接收者有两种形式 1. 值传递 2. 指针传递 如果你想改结构体的值 如果结构体的数据很大
	//go语言不支持继承 但是有办法能达到同样的效果 组合
}
