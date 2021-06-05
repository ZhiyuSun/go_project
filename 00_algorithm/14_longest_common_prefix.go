package main

// 14. 最长公共前缀
// 官方解法1 纵向
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// 官方解法2 横向法
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}


func main() {
	
}
