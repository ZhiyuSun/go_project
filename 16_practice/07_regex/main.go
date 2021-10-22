package main

import (
	"fmt"

	"regexp"
)

func main() {
	// https://c.runoob.com/front-end/854/
	str := "test asc"
	matched, err := regexp.MatchString("^\\S+\\s+(?:desc|asc|DESC|ASC)$", str)

	fmt.Println(matched, err)
}
