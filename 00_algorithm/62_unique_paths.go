package main

import "fmt"

// 62. 不同路径
// 2021.05.12 我自己写出来了，面对go语言，越来越游刃有余
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// 官方解法，在初始化值时更加的优雅
func uniquePaths1(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}



func main()  {
	fmt.Println(uniquePaths(3,7))
}