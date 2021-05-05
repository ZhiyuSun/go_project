package main

import (
	"context"
	"fmt"
	"go_project/11_grpc/05_grpc_test/proto"
	"google.golang.org/grpc"
)

func main() {
	//stream
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "sunzhiyu"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
