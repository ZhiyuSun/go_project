package main

import "fmt"

func replace(mySlice []string) {
	mySlice[0] = "sunny"
}

func main()  {
	// 什么是切片
	// 数组有一个很大的问题：大小确定，不能修改 - 切片 - 动态数组
	//var identifier []type
	//第一种方法： var courses []string //定义了一个切片
	var courses = []string{"django", "scrapy", "tornado"}
	fmt.Println(courses)
	fmt.Printf("%T\n", courses)

	//切片的第二种初始化方法 make
	//切片不是没有长度限制，为什么使用make初始化的需要我们传递一个长度， 那我传递了长度之后是不是意味着就像数组一样长度不能变了呢？
	courses1 := make([]string, 5)
	fmt.Println(len(courses1))
	//slice对标的就是python中list

	//第三种方法：通过数组变成一个切片
	var courses2 = [5]string{"django", "scrapy", "tornado","python", "asyncio"}
	subCourse := courses2[1:4] //亲切 //python中的用法叫切片 go语言中切片是一种数据结构 //python中切片的用法非常的多非常的灵活
	fmt.Println(subCourse)
	subCourse2 := courses2[:4] //亲切 //python中的用法叫切片 go语言中切片是一种数据结构 //python中切片的用法非常的多非常的灵活
	fmt.Println(subCourse2)
	// 在切片上操作会影响原来的数组
	subCourse[0] = "sunzhiyu"
	fmt.Println(subCourse)
	fmt.Println(courses2) // 原来的值也被改了

	//数组的传递是值传递 切片是引用传递
	//slice很重要，很常用
	replace(subCourse)
	fmt.Println(subCourse)
	fmt.Println(courses2)

	//第四种方式： new
	subCourse3 := new([]int)
	fmt.Println(subCourse3)

	//slice是动态数组，所以说我们需要动态添加值
	fmt.Println(subCourse[1])
	subCourse[1] = "rainy"
	fmt.Println(subCourse)

	subCourse4 := subCourse[1:4]
	fmt.Printf("%T, %v\n", subCourse, subCourse4)

	//append 可以向切片追加元素
	appendedCourse := []string{"imooc", "imooc2", "imooc3"}
	subCourse4 = append(subCourse4, appendedCourse...) //函数的参数传递规则
	fmt.Println(subCourse4)

	subCourse5 := make([]string, len(subCourse2))
	fmt.Println(len(subCourse5))
	copy(subCourse5,subCourse2)

	//append函数追加多个元素， 既然是切片，那为什么这个时候有来提出长度的要求
	//想从切片中删除元素怎么办？

	deleteCourses := [5]string{"django", "scrapy", "tornado","python", "asyncio"}
	courseSlice := deleteCourses[:]
	courseSlice = append(courseSlice[:1], courseSlice[2:]...) //取巧的做法
	fmt.Println(courseSlice)

	//如何判断某个元素是否在切片中

	//python和go语言的slice区别 玩出花样 go的slice更像是python的list，
	//go语言的底层是基于数组实现的 python的list底层也是基于数组实现的
	//slice进行的操作都会影响原来的数组， slice更像是一个指针 本身不存值

	// slice原理，因为很多底层的知识相对来说很多时候并不难而是需要花费比较多的时间去慢慢理解
	//1. 第一个现象
	a := make([]int, 0)
	b := []int{1,2,3}
	fmt.Println(copy(a, b)) // 0
	fmt.Println(a) // []，很奇怪
	//不会去扩展a的空间
	a1 := make([]int, 1)
	fmt.Println(copy(a1, b)) // 1
	fmt.Println(a1) // [1]

	//2. 第二个现象
	c := b[:]
	c[0] = 8
	fmt.Println(b) // [8 2 3] 影响了原数组
	fmt.Println(c) // [8 2 3]

	//3. 第三个现象
	c = append(c, 9)
	fmt.Println(b) //[8 2 3]， append函数没有影响到原来的数组
	fmt.Println(c) // [8 2 3 9]
	//这就是因为产生了扩容机制，扩容机制一旦产生 这个时候切片就会指向新的内存地址
	c[0] = 7
	fmt.Println(b) // [8 2 3]
	fmt.Println(c) // [7 2 3 9] 为什么append函数之后再调用c[0]=7不会影响到原来的数组

	//4. 第四个现象
	fmt.Println(c) // [7 2 3 9]
	fmt.Println(len(c)) // 4
	fmt.Println(cap(c)) // 6  cap指的是容量 长度和容量这两个概念
	fmt.Println(len(b)) // 3
	fmt.Println(cap(b)) // 3

	//切片底层是使用数组实现的，既要使用数组 又要满足动态的功能 怎么实现？
	//假设有一个值 实际上申请数组的时候可能是两个，如果后续要增加数据那么就直接添加到数据的结尾，这个时候我不要额外重新申请

	//切片有不同的初始化方式
	//1. 使用make方法初始化 len和cap是多少. 不会有多余的预留空间
	d := make([]int, 0)
	fmt.Printf("len=%d, cap=%d\n", len(d), cap(d)) // 0 0
	d = make([]int, 5)
	fmt.Printf("len=%d, cap=%d\n", len(d), cap(d)) // 5 5
	d = make([]int, 5, 6)
	fmt.Printf("len=%d, cap=%d\n", len(d), cap(d)) // 5 6

	//2. 通过数组取切片
	data := [10]int{0,1,2,3,4,5,6,7,8,9}
	slice := data[2:4] // 从2后面的值都属于slice
	newSlice := data[3:6]
	for index, value := range slice {
		fmt.Println(index, value)
	}
	fmt.Printf("len=%d, cap=%d\n", len(slice), cap(slice))
	fmt.Printf("len=%d, cap=%d\n", len(newSlice), cap(newSlice))

	//3.
	slice2 := []int{1,2,3}
	fmt.Printf("len=%d, cap=%d\n", len(slice2), cap(slice2)) // 3, 3

	//切片扩容问题， 扩容阶段会影响速度， python的list中底层实际上也是数组，也会面临动态扩容的问题，python的list中数据类型可以不一致
	oldSlice := make([]int, 0)
	fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice)) // 0, 0
	oldSlice = append(oldSlice,1)
	fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice)) // 1, 1
	oldSlice = append(oldSlice,2)
	fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice)) // 2, 2
	oldSlice = append(oldSlice,3)
	fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice)) // 3, 4
	oldSlice = append(oldSlice,4)
	fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice)) // 4, 4
	oldSlice = append(oldSlice,5)
	fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice)) // 5, 8
	oldSlice = append(oldSlice,6)
	oldSlice = append(oldSlice,7)
	oldSlice = append(oldSlice,8)
	fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice)) // 8, 8
	oldSlice = append(oldSlice,9)
	fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice)) // 9, 16

	/*
		Go 中切片扩容的策略是这样的：

		首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）
		否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap）
		否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的 1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
		如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）
	*/
	// 如果小于1024 扩容的速度是2倍 如果大于了1024 扩容的速度就是1.25
	oldSlice2 := make([]int, 1024)
	fmt.Printf("len=%d, cap=%d\n", len(oldSlice2), cap(oldSlice2)) // 1024, 1024
	oldSlice2 = append(oldSlice2,1)
	fmt.Printf("len=%d, cap=%d\n", len(oldSlice2), cap(oldSlice2)) // 1025, 1280

	//切片来说
	//1. 底层是数组，如果是基于数组产生的 会有一个问题就是会影响原来的数组。
	//2. 切片的扩容机制
	//3. 切片的传递是引用传递
	oldArr := [3]int{1,2,3}
	newArr := oldArr // 值传递
	newArr[0] = 5
	fmt.Println(newArr, oldArr) // [5 2 3] [1 2 3]

	oldSlice = []int{1,2,3}
	newSlice2 := oldSlice
	newSlice2[0] = 5
	fmt.Println(newSlice2, oldSlice) // [5 2 3] [5 2 3]

	//go语言中slice的原理讲解很重要，这个有坑（对于初学者），有经验的程序员觉得这个不是坑
	//程序员也是消费者-java c++ go -静态语言是站在对处理器的角度考虑， python - 才会站在使用者的角度上去考虑，对于处理器就不友好

	fmt.Println("===make中的坑===")
	//当make遇到了append容易出现的坑
	s1 := make([]int, 5) // 其实不是空的array
	fmt.Println(s1) // [0 0 0 0 0]
	s1 = append(s1, 6)
	//对于很多初学者来说我们期望的是只有一个数字就是6
	fmt.Println(s1) // [0 0 0 0 0 6]
	//很多人对make函数产生了一个假象 ，s1 := make([]int, 5) 好比是python的 s1 = []

	s2 := make([]int, 0)
	s2 = append(s2, 6)
	//对于很多初学者来说我们期望的是只有一个数字就是6
	fmt.Println(s2) // [6]
}
