package repository

import (
	"context"
)

func (m *Template[T]) Count(ctx context.Context, filter []string) (int64, error) {
	var count int64
	var entity T

	tx := MakeQuery(m.DB, filter)
	result := tx.Model(&entity).Count(&count)
	if err := result.Error; err != nil {
		return 0, err
	}

	return count, nil
}
