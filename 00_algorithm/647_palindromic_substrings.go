package main

import "fmt"

// 647. 回文子串

// 2021.05.26 参考了自己之前的动态规划法
func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i:=0; i<n; i++{
		dp[i] = make([]bool, n)
	}
	ans := 0
	for j:=0; j<n; j++{
		for i:=0; i<j+1; i++{
			if s[i] == s[j] && (j-i<2 || dp[i+1][j-1]){
				dp[i][j] = true
				ans ++
			}
		}
	}
	return ans
}

// 2021.05.26 官方解法，中心扩散法，这个范围定的很巧妙
func countSubstrings1(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < 2 * n - 1; i++ {
		l, r := i / 2, i / 2 + i % 2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings1("abc"))
}
