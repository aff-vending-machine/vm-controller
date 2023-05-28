package product

import (
	"context"

	"vm-controller/internal/layer/usecase/product/request"
	"vm-controller/pkg/helpers/db"

	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func (uc *usecaseImpl) Set(ctx context.Context, req *request.Set) error {
	if v := validate.Struct(req); !v.Validate() {
		return errors.Wrap(v.Errors.OneError(), "validate failed")
	}

	slots, err := uc.slotRepo.FindMany(ctx, db.NewQuery().AddWhere("product_sku = ?", req.Data.SKU))
	if err != nil {
		return err
	}

	for _, slot := range slots {
		if slot.Product != nil {
			json := map[string]interface{}{
				"product_name":      req.Data.Name,
				"product_group":     req.Data.Group,
				"product_image_url": req.Data.ImageURL,
				"product_price":     req.Data.Price,
			}

			uc.slotRepo.Update(ctx, db.NewQuery().AddWhere("code = ?", slot.Code), json)
		}
	}

	return nil
}
