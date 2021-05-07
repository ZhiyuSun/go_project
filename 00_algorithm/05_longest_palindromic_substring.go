package main

// 5. 最长回文子串

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		for l, r := i, i; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			if len(res) < r-l+1 {
				res = s[l : r+1]
			}
		}
		for l, r := i, i+1; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			if len(res) < r-l+1 {
				res = s[l : r+1]
			}
		}
	}
	return res
}


// 这个中心扩展算法的go语言版本很秀，可以利用指针去取代全局变量
func longestPalindrome1(s string) string {
	low, maxLen := 0, 0
	for i := range s {
		expand(s, i, i, &low, &maxLen)
		expand(s, i, i+1, &low, &maxLen)
	}
	return s[low : low+maxLen]
}
func expand(s string, l, r int, low, maxLen *int) {
	for ; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
	}
	if *maxLen < r-l-1 {
		*low = l + 1
		*maxLen = r - l - 1
	}
}

// 动态规划法
func longestPalindrome2(s string) string {
	lenth := len(s)

	if lenth <= 1 {
		return s
	}

	dp := make([][]bool, lenth)

	start := 0
	maxlen := 1

	for r := 0;r<lenth;r++ {
		dp[r] = make([]bool, lenth)
		dp[r][r] = true
		for l:=0;l<r;l++ {
			if s[l] == s[r] && (r -l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
			}else{
				dp[l][r] = false
			}

			if dp[l][r] {
				curlen := r-l+1
				if curlen > maxlen {
					maxlen = curlen
					start = l
				}
			}
		}
	}
	return s[start:start+maxlen]
}


func main()  {

}