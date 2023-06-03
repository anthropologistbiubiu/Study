package main

import (
	"context"
	"etcd/proto"
	"etcd/register"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8000, "The server port")
)

const testServerName = "s1"

// 定义一个server实现UnimplementedGreeterServer
// UnimplementedGreeterServer 是第四步自动生成的，可以打开对应文件查看
type server struct {
	proto.UnimplementedGreeterServer
}

// server 重写SayHello方法，做业务处理
func (s *server) SayHello(c context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	log.Printf("接收到客户端的消息: %v", req.GetName())
	time.Sleep(time.Second)
	ms := fmt.Sprintf("好的[%d]收到,%s %s", *port, req.GetName(), time.Now())
	log.Printf("回复客户端的消息: %s", ms)
	return &proto.HelloReply{Message: ms}, nil
}

func main() {
	// 解析命令行参数
	flag.Parse()
	// 监听本地tcp端口
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 创建一个grpc Server服务对象,Handler非必传
	// s := grpc.NewServer() // 可以直接创建对象
	s := grpc.NewServer()
	// 注册服务
	proto.RegisterGreeterServer(s, &server{})

	// 注册ETCD
	service, err := register.NewLocalDefNamingService("my1")
	if err != nil {
		log.Fatalf("failed to create NamingService: %v", err)
	}
	err = service.AddEndpoint(register.Endpoint{
		Addr:    "localhost",
		Name:    testServerName,
		Port:    *port,
		Version: "1.0.0",
	})
	if err != nil {
		log.Fatalf("failed to reg etcd: %v", err)
	}

	// 启动RPC并监听
	log.Printf("server listening at %v", lis.Addr())
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
			service.DelAllEndpoint()
		}
	}()

	// 等待关闭信号
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Printf("Shutdown Server ... \r\n")
	// 停止grpc服务
	s.GracefulStop()
	// 删除etcd注册信息
	service.DelAllEndpoint()
	fmt.Printf("Graceful Shutdown Server success\r\n")

}
