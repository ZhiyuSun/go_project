package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val int
	Next *ListNode
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {

	// 初始化slice的三种方法
	//var nilSlice []int
	//emptySlice1 := make([]int, 0)
	//emptySlice2 := []int{}

	// 排序方法
	nums := []int{1,3,2}
	sort.Ints(nums)
	fmt.Println(nums)

	// 二维数组的使用
	ans := make([][]int, 0)
	ans = append(ans, []int{1,2,3})
	fmt.Println(ans)

	// 定义map
	var phoneMap = map[string]string {
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	fmt.Println(phoneMap)

	// 空列表
	//var emptyStringList []string
	var emptyIntList []int
	emptyIntList = []int{}
	fmt.Println(emptyIntList)

	// 判断结构体为空？
	head := &ListNode{}
	fmt.Println(head)
	fmt.Printf("%T\n", head) // *main.ListNode, head如果是指针类型，可以跟nil做比较
	fmt.Printf("%T\n", *head) // main.ListNode

	// 如何在函数内部定义并调用一个递归函数
	// 参考22 括号生成

}