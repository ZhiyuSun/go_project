package main

import (
	"fmt"
	"sort"
)

// 49. 字母异位词分组

// 2021.05.11 排序法
func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

// 2021.05.11 计数法
func groupAnagrams2(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}


func main()  {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	oldStr := "sun"
	fmt.Println(oldStr)
	fmt.Printf("%T\n", oldStr)
	bytes := []byte(oldStr)
	fmt.Println(bytes)
	fmt.Printf("%T\n", bytes)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	fmt.Println(bytes)
	fmt.Printf("%T\n", bytes)
	newStr := string(bytes)
	fmt.Println(newStr)
	fmt.Printf("%T\n", newStr)

}
