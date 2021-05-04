package main

import "fmt"

func main()  {
	//和 C 语言的 for 一样：
	//for init; condition; post { }
	////和 C 的 while 一样：
	//for condition { }
	////和 C 的 for(;;) 一样：
	//for { }

	//for init; condition; post { } 计算1-10的和
	sum := 0
	for i := 1; i<=10; i++ {
		sum += i
	}
	fmt.Println(sum)
	//i只在for的局部变量之内，打印不出来
	//fmt.Println(i)

	// 类似while
	j := 0
	for ;j < 3; j ++ {
		fmt.Println("sun")
	}

	//循环字符串
	name := "name: 孙志宇"
	for index, value := range name {
		//为什么是数字 ， ascii
		//fmt.Println(index, value)
		fmt.Print(index)
		fmt.Printf("%c\n",  value)
	}

	fmt.Printf("%c\n", name[0])
	fmt.Printf("%c\n", name[9]) // 出现乱码

	//1. name是一个字符串， 2. 字符串是字符串的数组
	nameArr := []rune(name)
	for i := 0; i<len(nameArr); i++ {
		fmt.Printf("%c\n", nameArr[i]) // 不会出现乱码
	}

	//2. 在做字符串遍历的时候要尽量使用range

	// 其他例子
	strings := []string{"sun", "zhi", "yu"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i,x:= range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}

}
