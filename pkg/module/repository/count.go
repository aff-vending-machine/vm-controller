package repository

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/trace"
)

func (m *Template[T]) Count(ctx context.Context, filter []string) (int64, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	var count int64
	var entity T

	tx := MakeQuery(m.DB, filter)
	result := tx.Model(&entity).Count(&count)
	if err := result.Error; err != nil {
		return 0, err
	}

	return count, nil
}
