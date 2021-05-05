package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//1. 建立连接
	client, err := rpc.Dial("tcp", "localhost:2222")
	if err != nil {
		panic("连接失败")
	}
	var reply string //string有默认值
	err = client.Call("HelloService.Hello", "sunzhiyu", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)
}
