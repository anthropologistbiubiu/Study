package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	X, Y int
}

func main() {
	// 建立TCP连接
	fmt.Println("json client start")
	conn, err := net.Dial("tcp", "127.0.0.1:9091")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// 使用JSON协议
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	// 同步调用
	args := &Args{10, 20}
	var reply int
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal("ServiceA.Add error:", err)
	}
	fmt.Printf("ServiceA.Add: %d+%d=%d\n", args.X, args.Y, reply)

	// 异步调用
	var reply2 int
	divCall := client.Go("ServiceA.Add", args, &reply2, nil)
	replyCall := <-divCall.Done // 接收调用结果
	fmt.Println("replyCall.Error", replyCall.Error)
	fmt.Println("reply", reply2)
}
