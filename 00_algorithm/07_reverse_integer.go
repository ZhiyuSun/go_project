package main

import (
	"fmt"
	"math"
)

// 7. 整数反转

func reverse(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return
}


func main() {
	fmt.Println(reverse(10))
	fmt.Println(reverse(-12))
	fmt.Println(-12 % 10)
	fmt.Println(-12 / 10)
}
