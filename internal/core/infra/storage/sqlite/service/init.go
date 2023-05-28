package service

import (
	"context"
	"sync"

	"vm-controller/pkg/helpers/db"

	"gorm.io/gorm"
)

type Repository[T any] interface {
	Count(ctx context.Context, qry *db.Query) (int64, error)
	FindMany(ctx context.Context, qry *db.Query) ([]T, error)
	FindOne(ctx context.Context, qry *db.Query) (*T, error)
	Create(ctx context.Context, entity *T) (uint, error)
	Update(ctx context.Context, qry *db.Query, data map[string]interface{}) (int64, error)
	Delete(ctx context.Context, qry *db.Query) (int64, error)
}

type RepositoryImpl[T any] struct {
	db  *gorm.DB
	mtx sync.Mutex
}

func New[T any](db *gorm.DB) *RepositoryImpl[T] {
	return &RepositoryImpl[T]{
		db:  db,
		mtx: sync.Mutex{},
	}
}
