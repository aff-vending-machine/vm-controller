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
		s.ui.SendError(c.UserCtx, "identifiaction", "reference is mismatched")
		return fmt.Errorf("no reference in system: %s", reference)
	}

	if time.Since(s.stacks[reference].timestamp).Minutes() >= 5.0 {
		s.ui.SendError(c.UserCtx, "identifiaction", "OTP timeout")
		return fmt.Errorf("OTP timeout")
	}

	matched := otp == s.stacks[reference].otp && c.Data.Mail == s.stacks[reference].mail
	s.console_check(c, reference, otp, matched)
	if !matched {
		s.ui.SendError(c.UserCtx, "identifiaction", "OTP is mismatched")
		return fmt.Errorf("OTP is mismatched")
	}

	_, err := s.customerRepo.UpdateMany(
		c.UserCtx, []string{fmt.Sprintf("email:=:%s", c.Data.Mail)},
		map[string]interface{}{
			"cart":        c.Data.Raw(),
			"is_received": false,
		},
	)
	if err != nil {
		s.updateError(c, err)
		s.ui.SendEmergency(c.UserCtx, err)
		return errors.Wrap(err, "register failed")
	}

	s.updateIdentified(c)

	return nil
}
