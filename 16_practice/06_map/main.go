package main

import "fmt"

func main() {
	testMap := make(map[int]int, 3)
	testMap[1] = 1
	testMap[2] = 2
	testMap[3] = 3
	r, ok := testMap[2]
	fmt.Println(r)
	fmt.Println(ok)

	r2 := testMap[2]
	fmt.Println(r2)

	r3 := testMap[4]
	fmt.Println(r3)
}
