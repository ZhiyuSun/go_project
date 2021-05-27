package main

import "fmt"

// 739. 每日温度

// 2021.05.26 直奔题解，单调栈
func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	stack := []int{}
	for i := 0; i < length; i++ {
		temperature := temperatures[i]
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return ans
}


func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
