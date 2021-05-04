package main

import "fmt"

func main()  {
	//if语句
	num := 11
	if num % 2 == 0 {
		fmt.Println("偶数")
	}else{
		fmt.Println("奇数")
	}

	score := 55
	if score >= 90 {
		fmt.Println("优")
	}else if score >= 80 {
		fmt.Println("良")
	}else if score >= 60 {
		fmt.Println("一般")
	}else{
		fmt.Println("不及格")
	}

	// if statement; condition
	if num2 := 10; num2 % 2 == 0 {
		fmt.Println("偶数")
	}else{
		fmt.Println(num2)
	}
	// 不能打印，因为nums2不能出if else
	//fmt.Println(num2)

}
