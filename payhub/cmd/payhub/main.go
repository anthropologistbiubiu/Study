package main

import (
	"context"
	"flag"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/hashicorp/consul/api"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	tracing "payhub/internal/traceing"

	//"github.com/prometheus/client_golang/api"
	"os"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"payhub/internal/conf"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "payhub"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)
}

func newApp(logger log.Logger, gs *grpc.Server, hs1 *http.Server, hs2 *http.Server) *kratos.App {

	serviceRegistration1 := &api.AgentServiceRegistration{
		Name:    "payhub",
		ID:      "payhub-01",
		Port:    8000,
		Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			HTTP:     "http://127.0.0.1:8000/health",
			Method:   "GET",
			Header:   map[string][]string{"Content-Type": []string{"application/json"}},
			Interval: "5s", // 健康检查间隔时间
			Timeout:  "5s", // 超时时间
			//DeregisterCriticalServiceAfter: "1m", // 健康检查失败后取消注册
		},
		Weights: &api.AgentWeights{
			Passing: 10, // 设置为10
			Warning: 1,
		},
		// 其他配置项
	}
	serviceRegistration2 := &api.AgentServiceRegistration{
		Name:    "payhub",
		ID:      "payhub-02",
		Port:    8001,
		Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			HTTP:     "http://127.0.0.1:8001/health",
			Method:   "GET",
			Header:   map[string][]string{"Content-Type": []string{"application/json"}},
			Interval: "5s", // 健康检查间隔时间
			Timeout:  "5s", // 超时时间
			//DeregisterCriticalServiceAfter: "1m", // 健康检查失败后取消注册
		},
		Weights: &api.AgentWeights{
			Passing: 10, // 设置为10
			Warning: 1,
		},
		// 其他配置项
	}
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Agent().ServiceRegister(serviceRegistration1)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Agent().ServiceRegister(serviceRegistration2)
	if err != nil {
		log.Fatal(err)
	}
	reg := consul.New(client)

	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs1,
			hs2,
		),
		kratos.Registrar(reg),
	)
}

func main() {
	flag.Parse()
	/*
		logger := log.With(log.NewStdLogger(os.Stdout),
			"ts", log.DefaultTimestamp,
			"caller", log.DefaultCaller,
			"service.id", id,
			"service.name", Name,
			"service.version", Version,
			"trace.id", tracing.TraceID(),
			"span.id", tracing.SpanID(),
		)
	*/
	tp, err := tracing.NewTracerProvider("payhub-service-01", "http://localhost:14268/api/traces")
	if err != nil {
		log.Fatalf("Failed to initialize tracer provider: %v", err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Errorf("Error shutting down tracer provider: %v", err)
		}
	}()
	infoFile, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	errorFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer infoFile.Close()
	defer errorFile.Close()

	// 创建不同级别的日志处理器
	infoLogger := log.With(log.NewStdLogger(infoFile),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"level", log.LevelInfo)
	/*
		errorLogger := log.With(log.NewStdLogger(errorFile),
			"ts", log.DefaultTimestamp,
			"caller", log.DefaultCaller,
			"level", log.LevelError)
	*/
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, infoLogger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
