package main

import "fmt"

// 11. 盛最多水的容器
// 2021.05.08 我自己的方法
func maxArea(height []int) int {
	i, j, res := 0, len(height)-1, 0
	for i < j {
		if (j-i) * minValue(height[i], height[j]) > res {
			res =  (j-i) * minValue(height[i], height[j])
		}

		if height[i] < height[j] {
			i ++
		} else {
			j --
		}
	}
	return res
}

func minValue(a, b int) int{
	if a > b {
		return b
	} else {
		return a
	}
}

// 2021.05.08 别人的一个很有go味道的代码
func maxArea1(height []int) int {
	area, max := 0, 0
	left, right := 0, len(height)-1
	for left < right {
		if hl, hr, dist := height[left], height[right], right-left; hl > hr {
			area = dist * hr
			right--
		} else {
			area = dist * hl
			left++
		}

		if area > max {
			max = area
		}
	}
	return max
}


func main()  {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}