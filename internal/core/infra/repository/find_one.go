package repository

import (
	"context"
)

func (r *Template[T]) FindOne(ctx context.Context, filter []string) (*T, error) {
	var data T
	tx := MakeQuery(r.DB, filter)

	result := tx.First(&data)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &data, nil
}
