package order

import (
	"fmt"
)

func (s *stageImpl) reset() {
	s.step = 0
	s.slot = nil
	s.pendingItem.Clear()
}

func (s *stageImpl) nextStep() {
	s.step = s.step + 1
}

func (s *stageImpl) backStep() {
	s.step = s.step - 1
}

func makeSlotFilter(code string) []string {
	return []string{
		fmt.Sprintf("code:=:%s", code),
	}
}
