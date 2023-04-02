package receive

import (
	"github.com/aff-vending-machine/vm-controller/internal/core/domain/hardware"
	"github.com/aff-vending-machine/vm-controller/pkg/module/flow"
)

func (s *stageImpl) OnKeyPressed(c *flow.Ctx, key hardware.Key) error {
	switch s.status {
	case WAIT:
		if key.IsStar() {
			s.starcode = s.starcode + 1
			s.sharpcode = 0
		} else if key.IsSharp() {
			s.sharpcode = s.sharpcode + 1
			s.starcode = 0
		} else {
			s.sharpcode = 0
			s.starcode = 0
		}

	case CANCEL:
		s.queue.ClearStack(c.UserCtx)
		s.OnCancel(c, key)

	case DONE:
		s.queue.ClearStack(c.UserCtx)
		s.onDone(c, key)

	case E0:
		s.queue.ClearStack(c.UserCtx)
		s.onErrorE0(c, key)

	case E1:
		s.onErrorE1(c, key)

	case E2:
		s.queue.ClearStack(c.UserCtx)
		s.onErrorE0(c, key)

	default:
		return flow.ErrInvalidKey
	}

	if s.starcode == 6 {
		s.queue.ClearStack(c.UserCtx)
		s.OnCancel(c, key)
	}

	if s.sharpcode == 6 {
		s.queue.ClearStack(c.UserCtx)
		s.onDone(c, key)
	}

	return nil
}

func (s *stageImpl) onDone(c *flow.Ctx, key hardware.Key) error {
	switch key.Type() {
	case hardware.STAR:
		c.Reset()
		c.ChangeStage <- "idle"

	case hardware.SHARP:
		c.Reset()
		c.ChangeStage <- "order"
	}

	return nil
}

func (s *stageImpl) onErrorE0(c *flow.Ctx, key hardware.Key) error {
	switch key.Type() {
	case hardware.SHARP:
		c.Reset()
	}

	return nil
}

func (s *stageImpl) onErrorE1(c *flow.Ctx, key hardware.Key) error {
	switch key.Type() {
	case hardware.STAR:
		s.status = CANCEL
		s.displayUc.Background(c.UserCtx, "receive")
		s.show(c)

	case hardware.SHARP:
		s.status = 0
		s.queue.PushCommand(c.UserCtx, "COMMAND", "OPEN_GATE")
		s.displayUc.Background(c.UserCtx, "receive")
		s.show(c)
	}

	return nil
}

func (s *stageImpl) OnCancel(c *flow.Ctx, key hardware.Key) error {
	switch key.Type() {
	case hardware.STAR:
		c.Reset()
		c.ChangeStage <- "idle"

	case hardware.SHARP:
		c.Reset()
		c.ChangeStage <- "order"
	}

	return nil
}
