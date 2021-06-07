package main

import (
	"fmt"
	"math"
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

// 定义方向的方法
type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

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
	// 参考
	// 22. 括号生成
	// 39. 组合总和
	// 53. 全排列

	// 如何完成列表的拷贝
	var res [][]int
	path := []int{1,2,3}
	res = append(res, append([]int(nil), path...))
	path[0] = 4
	fmt.Println(res)
	fmt.Println(path)


	// 如何排序一个字符串
	oldStr := "sun"
	fmt.Println(oldStr)
	fmt.Printf("%T\n", oldStr)
	bytes := []byte(oldStr)
	fmt.Println(bytes)
	fmt.Printf("%T\n", bytes)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	fmt.Println(bytes)
	fmt.Printf("%T\n", bytes)
	newStr := string(bytes)
	fmt.Println(newStr)
	fmt.Printf("%T\n", newStr)

	// 如何声明较为复杂的数据结构
	//mp := map[string][]string{}
	//mp := map[[26]int][]string{}

	// map的键
	// golang中的map，的 key 可以是很多种类型，比如 bool, 数字，string, 指针, channel , 还有 只包含前面几个类型的 interface types, structs, arrays
	// 显然，slice， map 还有 function 是不可以了，因为这几个没法用 == 来判断
	mp := map[[3]int][]string{}
	mp[[...]int{1,2,3}] = []string{"123"}
	fmt.Println(mp)


	// 遍历字符串
	// 字符是uint8类型，可以直接相减
	rangeStr := "abcde"
	// 中文不知道怎么处理
	//rangeStr := "孙志宇"
	for i, v := range rangeStr {
		fmt.Println(i, v) // v会打印成ascll
		fmt.Printf("%T\n", v) // 是int32类型
		fmt.Println(v-'a')
	}

	// 排序专题
	intervals := [][]int{{15, 8}, {2, 6}, {1, 3}, {8, 10}}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)

	// 学习二维数组初始化，dp时很有用
	m, n := 4, 3
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// 最大值，最小值
	fmt.Println(math.MinInt64, math.MaxInt64)

	// 二叉树的层序遍历时，长度要用额外的变量表示，否则go里面会产生动态更新

	// 如果要比较三个数的大小
	maxV := max(1, max(2, 3))
	fmt.Println(maxV)

	// 哈希表的构造方法
	hashmap := map[int]int{}
	//hashmap1 := make(map[int]int)
	//var hashmap2 map[int]int
	if _, ok := hashmap[1]; ok {
		fmt.Println(1)
	}

	// 28. 实现 strStr()
	// 这里面有两个值得学的地方，一个是continue xxx，一个是遍历字符串用range
	aString := "sun"
	for i := range aString {
		fmt.Println(i)
	}

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}