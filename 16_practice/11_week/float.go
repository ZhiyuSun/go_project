package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var ff float64
	ff = 1.359123156
	ff = FloatRound(ff, 2)
	//fmt.Println(ff) // 输出 -1.3551

	//fmt.Println(fmt.Sprintf("%.1f", 3.0))
	//fmt.Println(fmt.Sprintf("%.1f", 3.1))
	//fmt.Println(fmt.Sprintf("%.1f", 3.2))
	//fmt.Println(fmt.Sprintf("%.1f", 3.51))
	fmt.Println(fmt.Sprintf("%.1f", 3.60-0.049))
	fmt.Println(fmt.Sprintf("%.1f", 3.61-0.049))
	fmt.Println(fmt.Sprintf("%.1f", 3.62-0.049))
	fmt.Println(fmt.Sprintf("%.1f", 3.63-0.049))
	fmt.Println(fmt.Sprintf("%.1f", 3.64-0.049))
	fmt.Println(fmt.Sprintf("%.1f", 3.65-0.049))
	fmt.Println(fmt.Sprintf("%.1f", 3.66-0.049))
	fmt.Println(fmt.Sprintf("%.1f", 3.67-0.049))
	fmt.Println(fmt.Sprintf("%.1f", 3.68-0.049))
	fmt.Println(fmt.Sprintf("%.1f", 3.69-0.049))
	fmt.Println(fmt.Sprintf("%.1f", 3.70-0.049))
	fmt.Println(time.Now().Weekday())
	now := time.Now()
	offset := int(time.Now().Weekday() - time.Sunday)
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, -offset)
	pName := "p" + weekStartDate.Format("20060102")
	fmt.Println(pName)
}

// 截取小数位数
func FloatRound(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}
