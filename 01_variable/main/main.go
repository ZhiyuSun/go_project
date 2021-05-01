package main

import "fmt"

func main()  {
	//变量的定义
	//静态的语言的变量和动态语言的变量定义差异很大
	//1. 最基础的变量定义
	var i int // int i
	i = 20
	fmt.Println(i)

	// 定义并初始化
	var j int = 10
	fmt.Println(j)

	// 2. 根据值自行判断变量类型（类型推断）
	var k = 100 // 并没有设置类型
	fmt.Println(k)

	// 3. go语言既然一门新语言 省略var
	m := 100
	fmt.Println(m)

	var a int = 10
	var b = 10
	c := 10
	fmt.Println(a, b , c)

	//多变量定义
	var d, e, f int
	d, e, f = 10, 20, 30
	fmt.Println(d, e, f)
	var p, q, r int = 10, 20 ,30
	fmt.Println(p, q, r)

	//// 集合类型
	var (
		age int
		name string
	)
	age = 1
	name = "sun"
	fmt.Println(age, name)

	// :=只能在初次使用时使用
	var s int =10
	s = 20 // s:= 20 会报错
	fmt.Println(s)

	//匿名变量， 变量一旦被定义 不使用，_ , 在go语言中会很常见

	//常量 -
	const PI = 3.1415926
	// 运行过程中， 代码写的不好 一不小心在某个地方将PI给改掉了
	// 不能修改const的值
	//PI = 3.14
	rr := 2.0
	cc := 2.0 * PI * rr
	fmt.Println(cc)

	// 枚举
	const (
		Unknown = 0
		Female = 1
		Male = 2
	)

	// 常量组如不指定类型和初始化值，该类型和值和上一行的类型一致
	const (
		ax int = 16
		ay
		as = "abc"
		az
	)
	fmt.Println(ax, ay, as, az)

	// 1. 常量的数据类型值可以是布尔，数字 和字符串
	// 2. 不曾使用的常量，在编译的时候是不会报错

	//const常量的iota， 常量计数器 枚举
	const (
		Book = iota //计算器
		Cloth //行
		Phone
		DeskTop
	)
	//0,1,2 本身不重要， 这三个值不一样
	//iota 该常量的值等于上一个常量的表达式
	fmt.Println(Book, Cloth, Phone, DeskTop)

	//iota你真的懂了吗？
	// 1. iota只能在常量组中是使用。iota:常量计数器
	// 2. 不同的const定义块互相不干扰
	const (
		U = iota
		F
		M
	)
	fmt.Println(U, F, M)

	// 3. 没有表达式的常量定义复用上一行的表达式
	// 4. 从第一行开始，iota从0逐行加一
	const (
		ta = iota //iota  = 0
		tb = 10 //iota = 1 每一行iota加一
		tc // c=10 , iota = 2
		td,te = iota, iota //iota=3
		tf = iota //iota=4
		//iota代表的是这里的行数, 不是写死的值，iota在运行中是会变化的
	)

	fmt.Println(ta, tb, tc, td, te, tf) // 0, 10, 10, 3, 3, 4

}
