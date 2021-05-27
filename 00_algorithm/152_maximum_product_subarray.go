package main

import "fmt"

// 152. 乘积最大子数组

// 2021.05.22 我自己的解法，很冗长
func maxProduct(nums []int) int {
	min := make([]int, len(nums))
	max := make([]int, len(nums))
	min[0] = nums[0]
	max[0] = nums[0]
	for i:=1;i<len(nums);i++{
		max[i] = maxValue(nums[i], nums[i] * min[i-1], nums[i] * max[i-1])
		min[i] = minValue(nums[i], nums[i] * min[i-1], nums[i] * max[i-1])
	}
	return maxValueOfArr(max)
}

func maxValue(a, b, c int)  int{
	if a < c {
		a = c
	}
	if a < b {
		a = b
	}
	return a
}

func minValue(a, b, c int)  int{
	if a > c {
		a = c
	}
	if a > b {
		a = b
	}
	return a
}

func maxValueOfArr(a []int)  int{
	res := a[0]
	for _,v := range a{
		if res < v {
			res = v
		}
	}
	return res
}

// 官方解法，省去了数组，同时max的思路很秀
func maxProduct2(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx * nums[i], max(nums[i], mn * nums[i]))
		minF = min(mn * nums[i], min(nums[i], mx * nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(maxProduct2([]int{2,3,-2,4}))
}
