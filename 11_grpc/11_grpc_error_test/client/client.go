package main

import (
	"context"
	"fmt"
	"go_project/11_grpc/11_grpc_error_test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"time"
)

func main() {
	//stream
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	// 测试超时，返回
	//context deadline exceeded
	//DeadlineExceeded
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)

	//go语言推荐的是返回一个error和一个正常的信息
	// grpc已经处理了语言之间的错误处理的不同方式
	// 超时机制，1. 网络抖动，2. 网络断开 A->B 10s
	_, err = c.SayHello(ctx, &proto.HelloRequest{Name: "sunzhiyu"})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			// Error was not a status error
			panic("解析error失败")
		}
		fmt.Println(st.Message())
		fmt.Println(st.Code())
	}
	//fmt.Println(r.Message)
}

