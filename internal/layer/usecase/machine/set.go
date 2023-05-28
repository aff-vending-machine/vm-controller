package machine

import (
	"context"

	"vm-controller/internal/layer/usecase/machine/request"
	"vm-controller/pkg/helpers/db"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (uc *usecaseImpl) Set(ctx context.Context, req *request.Set) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	updated := map[string]interface{}{
		"name":     req.Data.Name,
		"branch":   req.Data.Branch,
		"location": req.Data.Location,
		"type":     req.Data.Type,
		"vendor":   req.Data.Vendor,
		"status":   req.Data.Status,
	}

	_, err := uc.machineRepo.Update(ctx, db.NewQuery().AddWhere("id = ?", 1), updated)
	if err != nil {
		log.Error().Err(err).Interface("req", req).Msg("unable to update machine")
	}

	return nil
}
