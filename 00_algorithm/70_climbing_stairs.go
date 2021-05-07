package main

import "fmt"

// 我的解法
func climbStairs(n int) int {
	a := 0
	b := 1
	for i:=0; i<n; i++ {
		a, b = b, a + b
	}
	return b
}

// 官方解法
func climbStairs1(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func main()  {
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
}
