package main


// 27. 移除元素

// 2021.06.07 so easy
func removeElement(nums []int, val int) int {
	i, j := 0, 0
	for j<len(nums){
		if nums[j] != val {
			nums[i] = nums[j]
			i += 1
		}
		j += 1
	}
	return i
}


func main() {
	
}
