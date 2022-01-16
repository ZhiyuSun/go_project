package main

import (
	"fmt"
	"time"
)

func main() {
	var err error
	var t time.Time
	t, err = time.Parse("2006-01-02", "2021-09-01")
	if err != nil {
		fmt.Println("123")
	}
	end := t.AddDate(0, 0, 1)
	fmt.Println("end=", end)
	fmt.Printf("end=(%v)", end)
}
