package service

import (
	"context"

	"vm-controller/internal/core/infra/storage/sqlite"
	"vm-controller/pkg/helpers/db"
)

func (r *RepositoryImpl[T]) FindMany(ctx context.Context, query *db.Query) ([]T, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	var entities []T
	tx := sqlite.MakeQuery(r.db.WithContext(ctx), query)
	result := tx.Find(&entities)
	if err := result.Error; err != nil {
		return nil, err
	}
	return entities, nil
}
