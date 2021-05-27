package main

import "math"

// 581. 最短无序连续子数组

// 2021.05.26 直奔题解，很绕
func findUnsortedSubarray(nums []int) int {
	ret := 0
	left, right := -1, -1
	min, max := math.MaxInt64, math.MinInt64
	for i := 0; i < len(nums); i++ {
		//如果是递增数组，nums[i]>=max(nums[0,...,i-1])
		if nums[i] >= max {
			max = nums[i]
		} else {
			//持续更新最右
			right = i
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		//如果是递增数组，nums[i]<=min(nums[i,...,len])
		if nums[i] <= min {
			min = nums[i]
		} else {
			//持续更新最左
			left = i
		}
	}
	if right > left {
		ret = right - left + 1
	}
	return ret
}

func main() {
	
}
