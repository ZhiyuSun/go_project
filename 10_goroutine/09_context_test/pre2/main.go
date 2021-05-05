package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 2. 通过chan
// 刚才的两种方式并不统一
// go1.7 context
var stop chan bool = make(chan bool)
func cpuInfo(){
	defer wg.Done()
	for {
		// 有case和default,优先满足case
		select {
			case <- stop:
				fmt.Println("退出cpu监控")
				return
			default:
				time.Sleep(time.Second*2)
				fmt.Println("cpu信息读取完成")
		}
	}
}

func main(){
	// 现在启动一个goroutine去监控某台服务器的cpu使用情况
	wg.Add(1)
	go cpuInfo()
	// 现在希望可以中断cpu的信息读取

	time.Sleep(time.Second*6)
	stop <- true

	wg.Wait()
	fmt.Println("信息监控完成")



}
