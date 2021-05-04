package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 2. 通过chan
// 刚才的两种方式并不统一
// go1.7 context


func cpuInfo(ctx context.Context){
	defer wg.Done()
	ctx2, _ := context.WithCancel(ctx)
	go memInfo(ctx2)
	for {
		select {
			case <- ctx.Done():
				fmt.Println("退出cpu监控")
				return
			default:
				time.Sleep(time.Second*2)
				fmt.Println("cpu信息读取完成")
		}
	}
}

func memInfo(ctx context.Context){
	defer wg.Done()
	for {
		select {
		case <- ctx.Done():
			fmt.Println("退出内存监控")
			return
		default:
			time.Sleep(time.Second*2)
			fmt.Println("mem信息读取完成")
		}
	}
}

func main(){
	// 现在启动一个goroutine去监控某台服务器的cpu使用情况
	wg.Add(2)

	// 父context
	// 父context取消，父生成的子context也会取消
	ctx, cancel := context.WithCancel(context.Background())
	go cpuInfo(ctx)
	//go memInfo(ctx)
	time.Sleep(time.Second*6)

	cancel()
	// context在web开发中非常的常用，grpc每个接口调用都会传递context,gin的http接口也会有context
	wg.Wait()
	fmt.Println("信息监控完成")

}
