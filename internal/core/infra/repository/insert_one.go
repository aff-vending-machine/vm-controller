package repository

import (
	"context"
)

func (r *Template[T]) InsertOne(ctx context.Context, data *T) error {
	tx := r.DB.Begin()

	result := tx.Create(data)
	if err := result.Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
