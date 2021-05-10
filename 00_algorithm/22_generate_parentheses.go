package main

import "fmt"

// 22. 括号生成

// 2021.05.08 这种定义递归的方法也太巧妙了
func generateParenthesis(n int) []string {
	var res []string

	var dfs func(lRemain int, rRemain int, path string)
	dfs = func(lRemain int, rRemain int, path string) {
		if 2*n == len(path) {
			res = append(res, path)
			return
		}
		if lRemain > 0 {
			dfs(lRemain-1, rRemain, path+"(")
		}
		if lRemain < rRemain {
			dfs(lRemain, rRemain-1, path+")")
		}
	}

	dfs(n, n, "")
	return res
}


// 如何把函数放在外面，利用指针，我的蹩脚方法
func generateParenthesis2(n int) []string {
	var res []string
	dfs(n, n, "", &res)
	return res
}

// 这里传入的参数是数组的指针，另外我省去了传入的长度n
func dfs(lRemain int, rRemain int, path string, res *[]string) {
	if lRemain == 0 && rRemain == 0 {
		*res = append(*res, path)
		return
	}
	if lRemain > 0 {
		dfs(lRemain-1, rRemain, path+"(", res)
	}
	if lRemain < rRemain {
		dfs(lRemain, rRemain-1, path+")", res)
	}
}

func main()  {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis2(3))
}
