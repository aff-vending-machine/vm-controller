package repository

import (
	"context"
)

func (r *Template[T]) FindMany(ctx context.Context, filter []string) ([]T, error) {
	var entities []T
	tx := MakeQuery(r.DB, filter)

	result := tx.Find(&entities)
	if err := result.Error; err != nil {
		return nil, err
	}

	return entities, nil
}
