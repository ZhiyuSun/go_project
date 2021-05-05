package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	//返回值是通过修改reply的值
	*reply = "hello, " + request
	return nil
}

func main() {
	//1. 实例化一个server
	listener, _ := net.Listen("tcp", ":2222")
	//2. 注册处理逻辑 handler
	_ = rpc.RegisterName("HelloService", &HelloService{})
	//3. 启动服务
	for {
		conn, _ := listener.Accept() //当一个新的连接进来的时候，
		//rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) // 有弊端，多个请求过来一个一个处理
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
