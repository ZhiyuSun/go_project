package main

import (
	"context"
	"fmt"
	"go_project/11_grpc/05_grpc_test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	//stream
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	//md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	md := metadata.New(map[string]string{
		"name":    "szy",
		"pasword": "123456",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "sunzhiyu"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}

