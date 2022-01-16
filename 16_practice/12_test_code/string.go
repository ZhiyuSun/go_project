package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello,世界1234万"
	fmt.Println("方法一  格式化打印")
	for _, ch1 := range str {
		fmt.Printf("%q", ch1) //单引号围绕的字符字面值，由go语法安全的转义
	}
	fmt.Println("方法二  转化输出格式")
	for _, ch2 := range str {
		//fmt.Println(string(ch2))
		fmt.Println(string(ch2))
		fmt.Println(strings.Contains(str, string(ch2)))
	}
}
