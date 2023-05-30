package flow

type Stage string

const (
	CHANNEL_STAGE   = "channel"
	EMERGENCY_STAGE = "emergency"
	IDLE_STAGE      = "idle"
	ORDER_STAGE     = "order"
	PAYMENT_STAGE   = "payment"
	RECEIVE_STAGE   = "receive"
)
