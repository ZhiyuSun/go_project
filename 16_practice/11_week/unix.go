package main

import (
	"fmt"
	"time"
)

func main() {
	//tm := time.Unix(1638499284, 0)
	//fmt.Println(tm)
	//
	//round := 0
	//for range time.Tick(5 * time.Second) {
	//	fmt.Println("before", round)
	//	round += 1
	//	if round%2 == 0 {
	//		continue
	//	}
	//	fmt.Println("after", round)
	//}

	//targetValueFloat := float64(50)
	//targetValueStr := fmt.Sprintf("%.1f", targetValueFloat)
	//fmt.Println(targetValueStr)

	//midList := []int{1, 2, 3}
	//var mids []int
	//step := 2
	//for count := 0; count < len(midList); count += step {
	//	mids = midList[count : count+step]
	//	fmt.Println(mids)
	//}

	//var res []int
	//mids := []int{1, 2, 3}
	//mids2 := []int{}
	//res = append(res, mids...)
	//res = append(res, mids2...)
	//fmt.Println(res)

	count := 0
	for {
		time.Sleep(1 * time.Second)
		count += 1
		if count == 3 {
			//continue
			return
		}
		//fmt.Println("hello world")
		fmt.Println(count)
	}
}
