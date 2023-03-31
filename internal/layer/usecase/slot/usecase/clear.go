package slot_usecase

import "context"

func (uc *usecaseImpl) Clear(ctx context.Context, filter []string) error {
	_, err := uc.slotRepo.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
