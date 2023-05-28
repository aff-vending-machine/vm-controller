package repository

import (
	"context"
)

func (r *Template[T]) UpdateMany(ctx context.Context, filter []string, data map[string]interface{}) (int64, error) {
	var ent T
	tx := MakeQuery(r.DB.Begin(), filter)

	result := tx.Model(&ent).Updates(data)
	if err := result.Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()

	return result.RowsAffected, nil
}
