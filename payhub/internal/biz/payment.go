package biz

import (
	"context"
	"payhub/api/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
	Update(context.Context, *Greeter) (*Greeter, error)
	FindByID(context.Context, int64) (*Greeter, error)
	ListAll(context.Context) ([]*Greeter, error)
}

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}

type PaymentOrder struct {
	ID         int    `gorm:"primaryKey;autoIncrement"`
	MerchantID string `gorm:"column:merchantId;size:64;not null"`
	Amount     string `gorm:"size:64;not null"`
}
type PaymentRepo interface {
	Save(context.Context, *PaymentOrder) error
	CachePaymentOrder(context.Context, *PaymentOrder) error
}

type PaymentOrderUsecase struct {
	repo PaymentRepo
	log  *log.Helper
}

func NewPaymentOrderUsecase(repo PaymentRepo, logger log.Logger) *PaymentOrderUsecase {
	return &PaymentOrderUsecase{repo: repo, log: log.NewHelper(logger)}
}
func (uc *PaymentOrderUsecase) CreatePaymentOrder(ctx context.Context, g *PaymentOrder) error {
	uc.log.WithContext(ctx).Infof("CreatePaymentOrder,MerchantId: %v,Amount:%v",
		g.MerchantID, g.Amount)
	if err := uc.repo.Save(ctx, g); err != nil {
		return err
	}
	return uc.repo.CachePaymentOrder(ctx, g)
}
