package main

import "fmt"

// 55. 跳跃游戏

// 2021.05.11 参考Python的解法写了一下
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	endReachable := len(nums) - 1
	for i:=len(nums)-1; i>=0; i-- {
		if nums[i] + i >= endReachable {
			endReachable = i
		}
	}
	return endReachable == 0
}

func main()  {
	fmt.Println(canJump([]int{2,3,1,1,4}))
	fmt.Println(canJump([]int{3,2,1,0,4}))
}