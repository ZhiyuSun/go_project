package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := int64(100)
	b := int64(1)
	c := int64(0)
	d := float32(1)
	e := float32(12)
	f := float32(122)
	fmt.Println(float32(a / 10))
	fmt.Println(float32(b / 10))
	fmt.Println(float32(c / 10))
	fmt.Println(d / 10)
	fmt.Println(e / 10)
	fmt.Println(f / 10)

	k := float64(9900)
	strKm := strconv.FormatFloat(k, 'f', 1, 64)
	fmt.Println("StrKm = ", strKm)
}
