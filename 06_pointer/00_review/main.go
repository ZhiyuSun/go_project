package main

import "fmt"

func demo() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
}

func swap(a, b *int) {
	*b, *a = *a, *b
}

func swap2(a, b int) (int, int) {
	return b, a
}


func main() {
	demo()
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
	a, b = 3, 4
	a, b = swap2(a, b)
	fmt.Println(a, b)
}
