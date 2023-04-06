package slot

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
	"github.com/aff-vending-machine/vm-controller/pkg/trace"
)

func (r *repositoryImpl) Set(ctx context.Context, entiites []entity.Slot) error {
	_, span := trace.Start(ctx)
	defer span.End()

	tx := r.DB.Begin()
	tx = tx.Unscoped().Delete(entity.Slot{})
	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}

	tx = tx.Create(entiites)
	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
