package thaitropica

import "context"

func (uc *Flow) OnInit(ctx context.Context) {
	uc.context.UserCtx = ctx
	uc.stages[uc.context.Stage].OnInit(uc.context)
}
