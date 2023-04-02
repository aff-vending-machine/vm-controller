package identification

import (
	"fmt"
	"time"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/pkg/errors"
)

func (s *stageImpl) checkOTP(c *flow.Ctx, req WSReceived) error {
	reference := req.Data.Reference
	otp := req.Data.OTP

	if s.stacks[reference] == nil {
		err := fmt.Errorf("no reference in system: %s", reference)
		s.updateErrorTransaction(c, err)
		s.ui.SendError(c.UserCtx, "identifiaction", "reference is mismatched")
		return err
	}

	if time.Since(s.stacks[reference].timestamp).Minutes() >= 5.0 {
		err := fmt.Errorf("OTP timeout")
		s.updateErrorTransaction(c, err)
		s.ui.SendError(c.UserCtx, "identifiaction", "OTP timeout")
		return err
	}

	matched := otp == s.stacks[reference].otp && c.Data.Mail == s.stacks[reference].mail
	s.console_check(c, reference, otp, matched)
	if !matched {
		err := fmt.Errorf("OTP is mismatched")
		s.updateErrorTransaction(c, err)
		s.ui.SendError(c.UserCtx, "identifiaction", "OTP is mismatched")
		return err
	}

	_, err := s.customerRepo.UpdateMany(
		c.UserCtx, []string{fmt.Sprintf("email:=:%s", c.Data.Mail)},
		map[string]interface{}{
			"cart":        c.Data.Raw(),
			"is_received": false,
		},
	)
	if err != nil {
		s.updateErrorTransaction(c, err)
		s.ui.SendError(c.UserCtx, "identifiaction", "unable to register customer")
		c.ChangeStage <- "order"
		return errors.Wrap(err, "register failed")
	}

	s.updateIdentifiedTransaction(c)
	s.ui.SendIdentified(c.UserCtx, c.Data.MerchantOrderID, c.Data.TotalQuantity(), c.Data.TotalPrice())
	c.ChangeStage <- "receive"

	return nil
}
