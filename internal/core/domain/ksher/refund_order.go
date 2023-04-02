package ksher

type RefundOrderBody struct {
	RefundOrderID string  `json:"refund_order_id" validate:"required"` // Unlimited format, unique refund order number used by merchant
	Signature     string  `json:"signature" validate:"required"`       // Request signature
	Timestamp     string  `json:"timestamp" validate:"required"`       // Timestamp
	RefundAmount  float64 `json:"runfund_amount" validate:"required"`  // The refund amount, cents, cannot be greater than the order amount, only partial payment channels support non-full refunds.
	Provider      string  `json:"provider,omitempty"`                  // Payment gateway
	Mid           string  `json:"mid,omitempty"`                       // The merchant ID of this order is only valid when there are multiple sub-merchants under the same account.
}

type RefundOrderResult struct {
	ID              string `json:"id"`
	Status          string `json:"status"` // "Refunded"
	Timestamp       string `json:"timestamp"`
	Currency        string `json:"currency"`   // "THB"
	ErrorCode       string `json:"error_code"` // "REFUNDED"
	OrderType       string `json:"order_type"` //"Refund"
	Reference       string `json:"reference"`
	APIName         string `json:"api_name"`
	Locked          bool   `json:"locked"`
	Channel         string `json:"channel"` // wechat
	GatewayOrderID  string `json:"gateway_order_id"`
	Cleared         bool   `json:"cleared"`
	AcquirerOrderID string `json:"acquirer_order_id"`
	Amount          int    `json:"amount"`
	ErrorMessage    string `json:"error_message"`
	ChannelOrderID  string `json:"channel_order_id"`
	Signature       string `json:"signature"`
	ForceClear      bool   `json:"force_clear"`
	MerchantOrderID string `json:"merchant_order_id"`
	LogEntryURL     string `json:"log_entry_url"`
}
