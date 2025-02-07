package service

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	v12 "payhub/api/v1"
	"payhub/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v12.UnimplementedGreeterServer
	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v12.HelloRequest) (*v12.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v12.HelloReply{Message: "Hello " + g.Hello}, nil
}

type PaymentOrderService struct {
	v12.UnimplementedPaymentSerivceServer
	uc *biz.PaymentOrderUsecase
}

func NewPaymentOrderService(uc *biz.PaymentOrderUsecase) *PaymentOrderService {
	return &PaymentOrderService{uc: uc}
}

func (s *PaymentOrderService) CreatePaymentOrder(ctx context.Context, in *v12.PaymentCreateRequest) (*v12.PaymentCreateReply, error) {
	tracer := otel.Tracer("CreatePaymentOrder Service")
	// Start a new span
	ctx, span := tracer.Start(ctx, "CreatePaymentOrder")
	defer span.End()
	// Add attributes to the span
	span.SetAttributes(
		attribute.String("parameter", in.String()),
	)
	return &v12.PaymentCreateReply{
		Status: 200,
		PayUrl: "www.success.com",
	}, nil
	if err := s.uc.CreatePaymentOrder(ctx, &biz.PaymentOrder{MerchantID: in.Merchantid, Amount: in.Amount}); err != nil {
		return &v12.PaymentCreateReply{Status: 401, PayUrl: ""}, err
	}
	return &v12.PaymentCreateReply{Status: 200, PayUrl: "www.baidu.com"}, nil
}
func (s *PaymentOrderService) HealthCheck(ctx context.Context, in *v12.HealthRequest) (*v12.HealthReply, error) {
	return &v12.HealthReply{}, nil
}
