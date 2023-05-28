package repository

import (
	"context"
)

func (r *Template[T]) DeleteMany(ctx context.Context, filter []string) (int64, error) {
	var entity T
	tx := MakeQuery(r.DB.Begin(), filter)

	result := tx.Delete(&entity)
	if err := result.Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()

	return result.RowsAffected, nil
}
