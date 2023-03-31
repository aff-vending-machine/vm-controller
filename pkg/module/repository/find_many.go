package repository

import (
	"context"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/trace"
)

func (r *Template[T]) FindMany(ctx context.Context, filter []string) ([]T, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	var entities []T
	tx := MakeQuery(r.DB, filter)

	result := tx.Find(&entities)
	if err := result.Error; err != nil {
		return nil, err
	}

	return entities, nil
}
