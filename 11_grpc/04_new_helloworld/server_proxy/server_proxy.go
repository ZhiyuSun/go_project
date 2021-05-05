package server_proxy

import (
	"go_project/11_grpc/04_new_helloworld/handler"
	"net/rpc"
)

type HelloServicer interface {
	Hello(request string, reply *string) error
}

//如果做到解耦 - 我们关心的是函数 鸭子类型
func RegisterHelloService(srv HelloServicer) error {
	return rpc.RegisterName(hanlder.HelloServiceName, srv)
}
