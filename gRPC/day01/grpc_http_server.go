package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	X, Y int
}

type ServiceA struct{}

// Add 为ServiceA类型增加一个可导出的Add方法
func (s *ServiceA) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func main() {
	service := new(ServiceA)
	rpc.Register(service) // 注册RPC服务
	rpc.HandleHTTP()      // 基于HTTP协议
	l, e := net.Listen("tcp", ":9093")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	fmt.Println("server start")
	http.Serve(l, nil)
}
