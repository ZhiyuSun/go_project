package main

import (
	"fmt"

	"regexp"
)

func main() {
	// https://c.runoob.com/front-end/854/
	//str := "test asc"
	str := ""
	matched, err := regexp.MatchString("^\\S+\\s+(?:desc|asc|DESC|ASC)$", str)

	fmt.Println(matched, err)

	//var s []int
	//for i := 0; i < 3; i++ {
	//	s = append(s, i)
	//}
	//modifySlice(s)
	//fmt.Println(s)
	//
	//var a interface{}
	//fmt.Println(a == nil)
	//var b []int
	//fmt.Println(b == nil)
	//a = b
	//fmt.Println(a == nil)
}

func modifySlice(s []int) {
	s = append(s, 2048)
	//s = append(s, 4096)
	s[0] = 1024
}
