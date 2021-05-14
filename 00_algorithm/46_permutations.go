package main

import "fmt"

// 46. 全排列
// 2021.05.11 绝了，我自己写出来了
func permute(nums []int) [][]int {
	var res [][]int
	used := make([]int, len(nums))
	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			res = append(res, append([]int(nil), path...))
		}
		for i, v := range nums {
			if used[i] == 1 {
				continue
			}
			used[i] = 1
			dfs(append(path, v))
			used[i] = 0
		}
	}
	dfs([]int{})
	return res
}

func main()  {
	fmt.Println(permute([]int{1,2,3}))
}