package main

import (
	"fmt"
	"time"
)

// go语言诞生比较晚，web2.0开发逐渐主流，高并发
// 多线程-每个线程占用的内存比较多，而且系统切换开销很大，上千，绿城/轻量级线程-用户态的线程
// go语言一开始的时候就没有打算让我们去实例化一个线程 - 协程
// nodejs, python3.6也在支持，为什么go的协程会这么火，没有历史包袱
// python中有两种编程模式 1 多线程和多进程来进行并发编程 2 使用协程进行并发编程。很多库是基于多线程和多进程开发的。
// python，除非某一天所有的库，大部分库都支持了协程
// go的库都是用协程开发的

func p(){
	fmt.Println("sunzhiyu")
}

// 没有打印输出
// 因为程序有主线程
//func main()  {
//	go p()
//	// 主死从随
//	fmt.Println("hello") // 主线程能运行
//}

//func main()  {
//	go p() // 这是能打印了
//	time.Sleep(time.Second * 2)
//}


func p2() {
	for {
		fmt.Println(1)
	}
}
//
//func main()  {
//	for i:=0; i < 5; i++ {
//		go func() {
//			for {
//				fmt.Println(i) // 这个i随时变
//				time.Sleep(time.Second)
//			}
//		}()
//	} // 打印不规律，闭包特性
//	time.Sleep(time.Second*1)
//}


//func main() {
//	for i:=0; i < 1000; i++ {
//		go func(i int){
//			for {
//				fmt.Println(i)
//				time.Sleep(time.Second*1)
//			}
//		}(i)
//	}
//	time.Sleep(time.Second*1)
//}

// go的协程和Python的协程，网上有些人可能喜欢拿web框架来做性能对比
// go的gin begoo flask、django性能对比，这个不科学
// tornado sanic fastapi asyncio 协程库 我们就不再用python的多线程来对比
// 只要大家懂了任何一门语言的协程，其他语言的协程都很好理解 GMP
// 1. 大家都开启100w个协程
// 2. 使用的简单性， go语言的简单性

func main() {
	for i:=0; i < 1000000; i++ {
		go func(i int){
			for {
				fmt.Println(i)
				time.Sleep(time.Second)
			}
		}(i)
	}
	time.Sleep(time.Second*30)
}