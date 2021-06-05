package main

import "fmt"

// 26. 删除有序数组中的重复项
// 2021.06.05 自己摸索出来了，如果做不出来，我可以去si了
func removeDuplicates(nums []int) int {
	i, j := 0, 0
	for j<len(nums) {
		for j < len(nums)-1 && nums[j] == nums[j+1] {
			j += 1
		}
		nums[i] = nums[j]
		i += 1
		j += 1
	}
	return i
}

// 官方解法，反而有点难理解
func removeDuplicates2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// 参考之前python写过的比较清晰的思路
func removeDuplicates3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j:=1;j<len(nums);j++ {
		if nums[j] != nums[i] {
			i += 1
			nums[i] = nums[j]
		}
	}
	return i+1
}


func main() {
	fmt.Println(removeDuplicates([]int{1,1,2}))
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
	fmt.Println(removeDuplicates([]int{1}))
}
