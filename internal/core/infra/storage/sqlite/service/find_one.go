package service

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/infra/storage/sqlite"
	"github.com/aff-vending-machine/vm-controller/pkg/helpers/db"
)

func (r *RepositoryImpl[T]) FindOne(ctx context.Context, query *db.Query) (*T, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	var entity T
	tx := sqlite.MakeQuery(r.db.WithContext(ctx), query)
	result := tx.First(&entity)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &entity, nil
}
