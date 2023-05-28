package service

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/infra/storage/sqlite"
	"github.com/aff-vending-machine/vm-controller/pkg/helpers/db"
)

func (r *RepositoryImpl[T]) Count(ctx context.Context, query *db.Query) (int64, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	var count int64
	var entity T

	tx := sqlite.MakeQuery(r.db.WithContext(ctx), query)
	result := tx.Model(&entity).Count(&count)
	if err := result.Error; err != nil {
		return 0, err
	}

	return count, nil
}
