package main

import (
	"fmt"
	"time"
)

func main()  {
	/*
	go语言提供了一个select的功能，作用于channel之上的，多路复用
	select 会随机公平的选择一个case语句执行
	select的应用场景：
	1. timeout的超时机制
	2. 判断某个channel是否阻塞

	 */
	//ch1 := make(chan int, 1)
	//ch2 := make(chan int, 1)
	//ch1 <- 1
	//ch2 <- 2
	//select {
	//	case data := <- ch1:
	//		fmt.Println(data)
	//	case data := <- ch2:
	//		fmt.Println(data)
	//}

	//timeout := false
	//go func() {
	//	// 该goroutine执行时间超过了1s,那么就报告给主的groutine
	//	time.Sleep(time.Second*2)
	//	timeout = true
	//}()
	//
	//for {
	//	if timeout {
	//		fmt.Println("结束")
	//		break
	//	}
	//	time.Sleep(time.Millisecond*10)
	//}

	timeout := make(chan bool, 1)
	go func() {
		// 该goroutine执行时间超过了1s,那么就报告给主的groutine
		time.Sleep(time.Second*2)
		timeout <- true
	}()

	timeout2 := make(chan bool, 1)
	go func() {
		// 该goroutine执行时间超过了1s,那么就报告给主的groutine
		time.Sleep(time.Second*10)
		timeout2 <- true
	}()

	//fmt.Println(<- timeout)
	select {
		case <- timeout:
			fmt.Println("超时了")
		case <- timeout2:
			fmt.Println("超时了")
		default:
			fmt.Println("继续下一次") // 不想任何阻塞，设置default，防止阻塞住
	} // 不太懂default
}
