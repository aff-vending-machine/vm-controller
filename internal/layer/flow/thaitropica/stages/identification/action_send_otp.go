package identification

import (
	"fmt"
	"time"

	"github.com/aff-vending-machine/vmc-rpi-ctrl/internal/core/domain/mail"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/module/flow"
	"github.com/aff-vending-machine/vmc-rpi-ctrl/pkg/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (s *stageImpl) sendOTP(c *flow.Ctx, req WSReceived) error {
	email := req.Data.Mail
	reference := utils.GenerateRandom(6)
	otp := utils.GenerateOTP()
	timestamp := time.Now()
	err := s.mailAPI.Send(c.UserCtx, &mail.Message{
		From:    "thaitropica@gmail.com",
		To:      email,
		Subject: "Email verification for free sample TPK",
		Body: fmt.Sprintf(`
		<div>
      <p style="color:red">**THIS IS AUTOMATIC SYSTEM EMAIL. DO NOT REPLY.</p>
			<p>please verify your email by enter OTP code at vending machine.</p>
			<br/>
			<p><b>OTP: %s</b></p>
			<p><b>REF: %s</b></p>
			<br/>
			<p>Thank you</p>
      <p>For more information visit <a style="color:blue" href="https://www.thaitropica.co.th">www.thaitropica.co.th</a></p>
		</div>
		`, otp, reference),
	})
	if err != nil {
		log.Error().Err(err).Msg("unable to send mail")
		s.updateError(c, err)
		s.ui.SendEmergency(c.UserCtx, err)
		return errors.Wrap(err, "send Mail failed")
	}

	c.Data.Mail = email
	s.stacks[reference] = &OTPStack{
		mail:      email,
		reference: reference,
		otp:       otp,
		timestamp: timestamp,
	}
	s.console_request(c)

	log.Info().Str("OTP", otp).Str("Reference", reference).Msg("otp generate")
	err = s.ui.SendOTPRequest(c.UserCtx, c.Data.MerchantOrderID, email, reference, timestamp)
	if err != nil {
		s.updateError(c, err)
		s.ui.SendEmergency(c.UserCtx, err)
		return errors.Wrap(err, "send OPT requested to UI failed")
	}

	s.updateReference(c, email, reference, otp)

	return nil
}
