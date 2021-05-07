package main

import "fmt"

// 169. 多数元素
// 我的解法
func majorityElement(nums []int) int {
	var res, count int
	for _, v := range nums {
		if count == 0 {
			res = v
			count = 1
		} else {
			if res == v {
				count ++
			} else {
				count --
			}
		}
	}
	return res
}

// 参考别人的优化版
func majorityElement1(nums []int) int {
	major := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if major == num {
			count += 1
		} else {
			count -= 1
		}
	}

	return major
}


func main()  {
	fmt.Println(majorityElement([]int{3,2,3}))
}