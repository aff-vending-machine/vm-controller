package service

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/infra/storage/sqlite"
	"github.com/aff-vending-machine/vm-controller/pkg/helpers/db"
)

func (r *RepositoryImpl[T]) Update(ctx context.Context, query *db.Query, data map[string]interface{}) (int64, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	var entity T
	tx := sqlite.MakeQuery(r.db.WithContext(ctx).Begin(), query)
	result := tx.Model(&entity).Updates(data)
	if err := result.Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return result.RowsAffected, nil
}
