package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func consumer(msg chan int){
	defer wg.Done()
	fmt.Println(<-msg)
}

func main(){
	var msg chan int
	msg = make(chan int)

	//msg <-1
	//如果你是没有缓冲的channel 在没有启动一个消费者之前 你放数据就会报错
	//data := <- msg
	//fmt.Println(data)
	//当你进行 放数据到msg中的时候 这个时候会阻塞的，阻塞之前会获取一把锁， 这把锁什么时候释放 肯定是要等到数据被消费之后
	// 获取锁，等待消费者，但是没有等到，就死锁了
	// 就像快递员发快递，但是找不到你，就死锁了。但是如果有快递柜就不一样了。

	wg.Add(1)
	go consumer(msg)
	msg <- 1 //当你进行 放数据到msg中的时候 这个时候会阻塞的，阻塞之前会获取一把锁， 这把锁什么时候释放 肯定是要等到数据被消费之后
	wg.Wait()
	//channel是多个goroutine之间线程安全， 如何保证的呢 使用锁？

}
