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

	stop := false

	go func() {
		for {
			if stop {
				break
			}
			fmt.Println(1)
			time.Sleep(2 * time.Second) // 2 秒间隔
		}
	}()

	time.Sleep(10 * time.Second)
	//do_some_other_thing()
	stop = true
	fmt.Println(2)
	time.Sleep(10 * time.Second)
	fmt.Println(3)
	//for range time.Tick(2 * time.Second) {
	//	fmt.Println(2)
	//}

}

func exit1() {
	stop := false

	go func() {
		for {
			if !stop {
				break
			}
			//do_some_thing()
			time.Sleep(2 * time.Second) // 2 秒间隔
		}
	}()

	//do_some_other_thing()
	stop = false
}

func exit2() {
	t := time.NewTicker(time.Second * 2)

	go func() {
		for {
			select {
			case <-t.C:
				//do_some_thing()
			}
		}
	}()

	//do_some_other_thing()
	t.Stop()
}
