package main

import "fmt"

// 28. 实现 strStr()
// 2021.06.07 我的解法，能解决，但是蹩脚
func strStr(haystack string, needle string) int {
	if len(haystack) == 0 && len(needle) == 0 {
		return 0
	}

	for i:=0; i<len(haystack); i++ {
		if i + len(needle) > len(haystack) {
			break
		}
		result := true
		for j:=0; j<len(needle); j++ {
			if needle[j] != haystack[i+j] {
				result = false
				break
			}
		}
		if result {
			return i
		}
	}
	return -1
}


// 官方解法。这个outer秀到我了
func strStr2(haystack, needle string) int {
	n, m := len(haystack), len(needle)
outer:
	for i := 0; i+m <= n; i++ {
		// 这个range的写法也值得参考
		for j := range needle {
			if haystack[i+j] != needle[j] {
				continue outer
			}
		}
		return i
	}
	return -1
}


// 我参考官方后的优化版，感觉自己的代码在提升
func strStr3(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	for i:=0; i+m<=n; i++ {
		result := true
		for j := range needle {
			if needle[j] != haystack[i+j] {
				result = false
				break
			}
		}
		if result {
			return i
		}
	}
	return -1
}


func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("", ""))
}
