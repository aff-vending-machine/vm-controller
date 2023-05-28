package hardware

import (
	"strconv"
	"strings"

	"github.com/aff-vending-machine/vm-controller/pkg/helpers/gen"
	"github.com/rs/zerolog/log"
)

type Event struct {
	UID      string // 5
	Index    string // 1
	SlotCode string // 3
	Status   string // 2
}

type QueueHandler func(event *Event) error

func NewEvent(index int, item Item) Event {
	return Event{
		UID:      strings.ToUpper(gen.Random(5)),
		Index:    strconv.Itoa(index),
		SlotCode: item.SlotCode,
		Status:   "00",
	}
}

func NewEventFromString(code string) *Event {
	if len(code) != 11 {
		log.Error().Str("code", code).Int("size", len(code)).Msg("invalid event code")
		return nil
	}

	return &Event{
		UID:      code[0:5],
		Index:    code[5:6],
		SlotCode: code[6:9],
		Status:   code[9:11],
	}
}

func (e *Event) ToValueCode() string {
	return e.UID + e.Index + e.SlotCode + e.Status
}
