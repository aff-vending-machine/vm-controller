package identification

import (
	"fmt"
	"strings"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/entity"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/pkg/errors"
)

func (s *stageImpl) requestOTP(c *flow.Ctx, req WSReceived) error {
	var customer *entity.Customer
	var err error
	email := req.Data.Mail

	customer, err = s.customerRepo.FindOne(c.UserCtx, []string{fmt.Sprintf("email:=:%s", email)})
	if err != nil && strings.Contains(err.Error(), "not found") {
		customer = &entity.Customer{Email: email}
		err := s.customerRepo.InsertOne(c.UserCtx, customer)
		if err != nil {
			s.updateError(c, err)
			s.ui.SendEmergency(c.UserCtx, err)
			return errors.Wrap(err, "unable to create customer")
		}
	} else if err != nil {
		s.updateError(c, err)
		s.ui.SendEmergency(c.UserCtx, err)
		return errors.Wrap(err, "register failed")
	}

	if customer.IsReceived {
		s.console_email_used(c, email)
		s.ui.SendMailIsUsed(c.UserCtx, c.Data.MerchantOrderID, email)
		return errors.Wrap(err, "this email is used")
	}

	return s.sendOTP(c, req)
}
