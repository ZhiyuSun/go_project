package main

import (
	"context"
	"fmt"
	"go_project/11_grpc/07_grpc_proto_test/proto_bak" // 根据proto文件自动生成的代码
	"google.golang.org/grpc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

func main() {
	// 创建连接
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Printf("连接失败: [%v]\n", err)
		return
	}
	defer conn.Close()
	// 声明客户端
	client := proto_bak.NewGreeterClient(conn)
	rsp, _ := client.SayHello(context.Background(), &proto_bak.HelloRequest{
		Name: "sunzhiyu",
		Url:  "https://baidu.com",
		G:    proto_bak.Gender_MALE,
		Mp: map[string]string{
			"name":    "sunzhiyu",
			"company": "coding",
		},
		AddTime: timestamppb.New(time.Now()),
	})
	fmt.Println(rsp.Message)
}
