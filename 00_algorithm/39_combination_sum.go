package main

import "fmt"

// 39. 组合总和

// 2021.05.10 尝试自己把这道题给做出来
func combinationSum(candidates []int, target int) [][]int {
	var res = make([][]int, 0)
	// 上面的另一种写法
	// var res [][]int
	dfs(0, target, &res, []int{}, candidates)
	return res
}


func dfs(index int, rest int, res *[][]int, path []int, candidates []int) {
	if rest == 0 {
		*res = append(*res, append([]int(nil), path...)) // 我的方法，问题出在这里，后来调整完没问题了
		return
	}
	for i:=index;i<len(candidates);i++{
		if candidates[i] <= rest {
			path = append(path, candidates[i])
			dfs(i, rest-candidates[i], res, path, candidates)
			path = path[:len(path)-1]
		}
	}
}


// 2021.05.10 如果把dfs放在外面，会增加很多的变量，反而把程序变得复杂
// 所以下次遇到递归回溯的题目，就在主函数里定义dfs函数
func combinationSum2(candidates []int, target int) [][]int {
	var ans [][]int
	var comb []int
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...)) // 这个复制的方法，可以学一下
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1] // 去掉末尾元素的做法，跟我想的一样
		}
	}
	dfs(target, 0)
	return ans
}



func main()  {
	fmt.Println(combinationSum([]int{2,3,6,7}, 7))
}
