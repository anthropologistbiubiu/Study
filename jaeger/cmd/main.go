package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func HttpPost(ctx context.Context, url string, bodyStr string, timeout int, header map[string]string) ([]byte, error) {
	reader := strings.NewReader(bodyStr)
	request, err := http.NewRequestWithContext(ctx, "POST", url, reader)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		request.Header.Set(k, v)
	}
	client := &http.Client{
		Timeout:   time.Second * time.Duration(timeout),
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBytes, nil
}

func Sign(ctx context.Context, url string, bodyStr string) []byte {

	sign := md5.Sum([]byte(url + bodyStr))
	return sign[:]

}

func init() {
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)
}
func main() {
	// 创建 Jaeger 导出器
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	if err != nil {
		panic(fmt.Sprintf("failed to initialize exporter: %v", err))
	}

	// 创建 TracerProvider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("gin-service-gateway"),
		)),
	)
	defer func() { _ = tp.Shutdown(context.Background()) }()

	// 设置全局 TracerProvider
	otel.SetTracerProvider(tp)

	route := gin.Default()
	route.POST("/test/jaeger", func(c *gin.Context) {

		bodyJson := `{"merchantid":"10001"}`
		url := "http://localhost:8000/payment/create"
		tracer := otel.Tracer("jager-service-01")
		ctx, span := tracer.Start(context.Background(), "jaeger-service-sign")
		defer span.End()

		Sign(ctx, url, bodyJson)

		header := map[string]string{
			"Content-Type": "application/json",
		}
		ctx, span = tracer.Start(context.Background(), "jaeger-service-http")
		defer span.End()
		respBytes, err := HttpPost(ctx, url, bodyJson, 5, header)
		if err != nil {
			fmt.Println("HttpPost Err", err)
		}
		c.String(http.StatusOK, string(respBytes))

	})
	route.Run(":8003")
}
