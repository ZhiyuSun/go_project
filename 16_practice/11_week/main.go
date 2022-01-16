package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(time.Monday)
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	res := weekStartDate.Format("2006-01-02")
	fmt.Println(res)
	fmt.Println(time.Now())

	a := []int{1, 2, 3}
	fmt.Println(a[:1])

	f := 3.1415926

	s := strconv.FormatFloat(f, 'f', -1, 64)
	fmt.Println(s)

	fmt.Println(fmt.Sprintf("%.1f", 3.0))
	v, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", 3.61), 64)
	fmt.Println(v)

	l := 156000
	fmt.Println(float64(l) / 10000)

}
