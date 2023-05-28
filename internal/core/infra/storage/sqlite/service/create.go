package service

import (
	"context"
)

func (r *RepositoryImpl[T]) Create(ctx context.Context, entity *T) (uint, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	tx := r.db.WithContext(ctx).Begin()
	result := tx.Create(entity)
	if err := result.Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return uint(result.RowsAffected), nil
}
