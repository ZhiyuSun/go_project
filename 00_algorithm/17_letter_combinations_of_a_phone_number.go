package main

import "fmt"

// 17. 电话号码的字母组合

// 2021.05.08 经典回溯题，决定抄一遍官方解法
// 这道题值得好好学习
var phoneMap = map[string]string {
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return combinations
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string)  {
	if index == len(digits){
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index]) // uint8 to string
		letters := phoneMap[digit]
		lettersCount := len(letters)
		for i:=0; i<lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}

func main()  {
	fmt.Println(letterCombinations("23"))
}