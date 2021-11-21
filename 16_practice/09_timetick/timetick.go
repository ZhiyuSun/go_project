package main

import (
	"fmt"
	"time"
)

func main() {
	//for {
	//	time.Sleep(2 * time.Second) // 2 秒间隔
	//	fmt.Println(1)
	//}

	t := time.NewTicker(time.Second * 2)

	go func() {
		for {
			select {
			case <-t.C:
				fmt.Println(1)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println(2)
	t.Stop()
	time.Sleep(10 * time.Second)
	fmt.Println(3)
}
