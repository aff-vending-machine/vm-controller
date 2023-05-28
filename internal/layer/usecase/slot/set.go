package slot

import (
	"context"

	"github.com/aff-vending-machine/vm-controller/internal/layer/usecase/slot/request"
	"github.com/aff-vending-machine/vm-controller/pkg/helpers/db"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Set(ctx context.Context, req *request.Set) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	slots, err := uc.slotRepo.FindMany(ctx, db.NewQuery())
	if err != nil {
		return err
	}

	mapCase := make(map[string]int)
	mapIndex := make(map[string]int)

	for index, cs := range req.Data {
		mapCase[cs.Code] += 1
		mapIndex[cs.Code] = index
	}

	for _, ms := range slots {
		mapCase[ms.Code] += 2
	}

	counts := []int{0, 0, 0, 0}
	for code, condition := range mapCase {
		switch condition {
		case 1: // only in center, add it
			counts[0]++
			slot := req.Data[mapIndex[code]]
			uc.slotRepo.Create(ctx, slot.ToEntity())

		case 2: // only in machine, remove it
			counts[1]++
			uc.slotRepo.Delete(ctx, db.NewQuery().AddWhere("code = ?", code))

		case 3: // both, update it
			counts[2]++
			slot := req.Data[mapIndex[code]]
			uc.slotRepo.Update(ctx, db.NewQuery().AddWhere("code = ?", code), slot.ToJson())

		default:
			counts[3]++
			log.Warn().Int("condition", condition).Str("Code", code).Interface("data", req.Data[mapIndex[code]]).Msg("invalid condition")
		}
	}
	log.Info().Int("total", len(mapCase)).Int("add", counts[0]).Int("remove", counts[1]).Int("update", counts[2]).Int("ignore", counts[3]).Msg("sync slot data")

	return nil
}
