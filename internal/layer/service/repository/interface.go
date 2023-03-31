package repository

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/entity"
)

type Customer interface {
	Count(ctx context.Context, filter []string) (int64, error)
	FindOne(ctx context.Context, filter []string) (*entity.Customer, error)
	FindMany(ctx context.Context, filter []string) ([]entity.Customer, error)
	InsertOne(ctx context.Context, ent *entity.Customer) error
	UpdateMany(ctx context.Context, filter []string, ent map[string]interface{}) (int64, error)
	DeleteMany(ctx context.Context, filter []string) (int64, error)
}

type Machine interface {
	Count(ctx context.Context, filter []string) (int64, error)
	FindOne(ctx context.Context, filter []string) (*entity.Machine, error)
	FindMany(ctx context.Context, filter []string) ([]entity.Machine, error)
	InsertOne(ctx context.Context, ent *entity.Machine) error
	UpdateMany(ctx context.Context, filter []string, ent map[string]interface{}) (int64, error)
	DeleteMany(ctx context.Context, filter []string) (int64, error)
}

type PaymentChannel interface {
	Count(ctx context.Context, filter []string) (int64, error)
	FindOne(ctx context.Context, filter []string) (*entity.PaymentChannel, error)
	FindMany(ctx context.Context, filter []string) ([]entity.PaymentChannel, error)
	InsertOne(ctx context.Context, ent *entity.PaymentChannel) error
	UpdateMany(ctx context.Context, filter []string, ent map[string]interface{}) (int64, error)
	DeleteMany(ctx context.Context, filter []string) (int64, error)
}

type Slot interface {
	Count(ctx context.Context, filter []string) (int64, error)
	FindOne(ctx context.Context, filter []string) (*entity.Slot, error)
	FindMany(ctx context.Context, filter []string) ([]entity.Slot, error)
	InsertOne(ctx context.Context, ent *entity.Slot) error
	UpdateMany(ctx context.Context, filter []string, ent map[string]interface{}) (int64, error)
	DeleteMany(ctx context.Context, filter []string) (int64, error)
}

type Transaction interface {
	Count(ctx context.Context, filter []string) (int64, error)
	FindOne(ctx context.Context, filter []string) (*entity.Transaction, error)
	FindMany(ctx context.Context, filter []string) ([]entity.Transaction, error)
	InsertOne(ctx context.Context, ent *entity.Transaction) error
	UpdateMany(ctx context.Context, filter []string, ent map[string]interface{}) (int64, error)
	DeleteMany(ctx context.Context, filter []string) (int64, error)
}
