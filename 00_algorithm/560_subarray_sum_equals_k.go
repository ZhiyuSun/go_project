package main

import "fmt"

// 560. 和为K的子数组
// 2021.05.27 我自己的做法
func subarraySum(nums []int, k int) int {
	s, res := 0, 0
	hashmap := make(map[int]int)
	hashmap[0] = 1
	for _,v:= range nums {
		s += v
		value, ok := hashmap[s-k]
		if ok {
			res += value
		}
		hashmap[s] += 1
	}
	return res
}

// 其他解法赏析，枚举
func subarraySum1(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end >= 0; end-- {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}


// 前缀和+哈希表
func subarraySum2(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre - k]; ok {
			count += m[pre - k]
		}
		m[pre] += 1
	}
	return count
}


func main() {
	fmt.Println(subarraySum([]int{1,1,1}, 2))
}
