package main

import (
	"fmt"
	"strconv"
)

func main()  {
	//printf 格式化
	name := "sunzhiyu"
	age := 26
	fmt.Println("name:"+name+", age:"+strconv.Itoa(age))
	fmt.Printf("name:%T, age: %T\n", name, age)
	fmt.Printf("name:%s, age:%x,\n ", name, age) // 16进制
	desc := fmt.Sprintf("name:%s, age:%x,\n ", name, age) // 变成字符串
	fmt.Println(desc)

	data := 65
	fmt.Printf("%c\n", data)
	fmt.Printf("%q\n", data)
	fmt.Printf("%e", 65.1)

	//输入
	var n string
	var a int
	fmt.Println("请输入你的姓名和年龄:")
	fmt.Scanln(&n, &a)
	fmt.Println(n, a)

}
