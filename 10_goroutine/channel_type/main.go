package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
func consumer(queue chan int) {
	defer wg.Done()
	for{
		data, ok := <- queue
		if !ok {
			break
		}
		fmt.Println(data, ok)
		time.Sleep(time.Second)
		queue <- 2
	}

}

func main()  {
	// 有没有缓冲，1 有缓冲 2 无缓冲
	// 双向的还是单向的

	var msg chan int // 静态语言要定义类型
	msg = make(chan int) // 第一种初始化方式，无缓冲
	wg.Add(1)
	go consumer(msg) // 普通的channel可以直接转换为单向的channel
	msg <- 1 // 将1放入到channel中
	fmt.Println("等待返回值")
	fmt.Println(<-msg)
	close(msg)
	wg.Wait()

}
