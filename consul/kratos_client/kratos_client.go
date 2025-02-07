package main

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/filter"
	"github.com/go-kratos/kratos/v2/selector/random"
	"github.com/go-kratos/kratos/v2/selector/wrr"
	kratos_http "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/hashicorp/consul/api"
	"io"
	"net/http"
	"strings"
)

func main() {
	// 创建 Consul 客户端
	consulClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	// 使用 Consul 作为服务发现
	r := consul.New(consulClient)
	// 创建路由 Filter：筛选版本号为"2.0.0"的实例
	versionFilter := filter.Version("1.0.0")
	// 设置全局的 Selector，使用 wrr 算法
	selector.SetGlobalSelector(wrr.NewBuilder())

	selector.SetGlobalSelector(random.NewBuilder())
	// 创建 HTTP 客户端
	hConn, err := kratos_http.NewClient(
		context.Background(),
		kratos_http.WithEndpoint("discovery:///payhub"),
		kratos_http.WithDiscovery(r),
		kratos_http.WithNodeFilter(versionFilter),
		kratos_http.WithBlock(),
	)
	fmt.Printf("&&&&&&&&&&&&&&&&&& %+v \n", hConn)
	if err != nil {
		panic(err)
	}
	// 使用客户端发送请求
	reqBody := `a"merchantId": "100001"}` // 示例请求体
	req, err := http.NewRequest("POST", "/payment/create", strings.NewReader(reqBody))
	if err != nil {
		fmt.Printf("Failed to create request: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	response, err := hConn.Do(req)
	if err != nil {
		fmt.Printf("Failed to make request: %v,%v\n", err, response)
		return
	}
	if resBytes, err := io.ReadAll(response.Body); err != nil {
		fmt.Println("///////////", err, resBytes)
	}
}
