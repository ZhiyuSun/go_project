package main

import "fmt"

// 33. 搜索旋转排序数组
// 2021.05.10 参照了Python，自己写了go的方法
func search(nums []int, target int) int {
	if len(nums) == 0{
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l+r)/2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[len(nums)-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func main()  {
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 0))
}