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
	}

}

func main()  {
   /*
	channel提供了一种通信机制，定下，python,java消息队列
    */
	// 1. 定义一个channel
	var msg chan int // 静态语言要定义类型
	//var a *int
	//if msg == nil{ // 引用类型不初始化默认是nil类型
	//
	//}

	// 2. 初始化这个channel，两种方式
	//msg = make(chan int) // 第一种初始化方式，无缓冲
	msg = make(chan int, 1) // 第二种初始化方式：有缓冲空间
	// 3. 在go语言中，使用make初始化的有3中，slice, map, channel
	msg <- 1 // 将1放入到channel中
	//msg <- 2 // 死锁
	wg.Add(1)
	go consumer(msg)
	msg <- 2 // 再放，没有问题
	// 关闭channel
	// 1. 已经关闭的channel不能再发送数据了
	// 2. 已经关闭的channel，消费者能继续取数据，直到数据取完
	close(msg)
	wg.Wait()
	//msg <- 2 // 你这个管道看起来好像就是一个有空间的数组
	//data := <- msg // 将箭头右边的值放入到左边
	//fmt.Println(data)


}
