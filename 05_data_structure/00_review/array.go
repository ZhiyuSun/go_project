package main

import "fmt"

func printArray(arr [5]int){
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray2(arr *[5]int){
	arr[0] = 101
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray3(arr []int){
	arr[0] = 103
	for i, v := range arr {
		fmt.Println(i, v)
	}
}


func main() {
	var arr1 [5]int
	arr2 := [3]int{1,3,5}
	arr3 := [...]int{2,4,6,8,10}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := range arr3 {
		fmt.Println(i)
	}
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	fmt.Println("printArray(arr1)")
	printArray(arr1)

	fmt.Println("printArray(arr3)")
	printArray(arr3)

	fmt.Println("printArray(arr1, arr3)")
	fmt.Println(arr1, arr3)

	fmt.Println("printArray(arr1)")
	printArray2(&arr1)

	fmt.Println("printArray(arr3)")
	printArray2(&arr3)

	fmt.Println("printArray(arr1, arr3)")
	fmt.Println(arr1, arr3)

	fmt.Println("printArray(arr1)")
	printArray3(arr1[:])

	fmt.Println("printArray(arr3)")
	printArray3(arr3[:])

	fmt.Println("printArray(arr1, arr3)")
	fmt.Println(arr1, arr3)


}
