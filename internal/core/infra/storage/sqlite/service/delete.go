package service

import (
	"context"

	"vm-controller/internal/core/infra/storage/sqlite"
	"vm-controller/pkg/helpers/db"
)

func (r *RepositoryImpl[T]) Delete(ctx context.Context, query *db.Query) (int64, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	var entity T
	tx := sqlite.MakeQuery(r.db.WithContext(ctx).Begin(), query)
	result := tx.Delete(&entity)
	if err := result.Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return result.RowsAffected, nil
}
