package main

import "fmt"

// 283. 移动0
// 我的解法
func moveZeroes(nums []int)  {
	i := 0
	for j, v := range nums {
		if v != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i ++
		}
	}
}

// 官方解法
func moveZeroes1(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func main()  {
	arr := []int{0,1,0,3,12}
	moveZeroes(arr)
	fmt.Println(arr)
}