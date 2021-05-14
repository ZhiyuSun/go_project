package main

import "fmt"

// 78. 子集
// 2021.05.12 我的解法，磕磕绊绊，写出来了
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	var dfs func(index int, path []int)
	dfs = func(index int, path []int) {
		if index == len(nums) {
			res = append(res, append([]int(nil), path...))
			return
		}
		path = append(path, nums[index])
		dfs(index+1, path)
		path = path[:len(path)-1]
		dfs(index+1, path)
	}
	dfs(0, []int{})
	return res
}


// 官方解法赏析
func subsets1(nums []int) (ans [][]int) {
	var set []int
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}

func main()  {
	fmt.Println(subsets([]int{1,2,3}))
}
