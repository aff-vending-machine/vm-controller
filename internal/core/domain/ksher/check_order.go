package ksher

type CheckOrderQuery struct {
	MID       string `json:"mid,omitempty"`                 // The merchant ID of this order is only valid when there are multiple sub-merchants under the same account.
	Provider  string `json:"provider,omitempty"`            // Payment gateway
	Signature string `json:"signature" validate:"required"` // Request signature
	Timestamp string `json:"timestamp" validate:"required"` // Timestamp
}

type CheckOrderResult struct {
	Cleared         bool   `json:"cleared"`
	Channel         string `json:"channel"` // wechat
	Status          string `json:"status"`  // "Refunded", "Closed"
	Note            string `json:"note"`
	MerchantOrderID string `json:"merchant_order_id"`
	Signature       string `json:"signature"`
	Timestamp       string `json:"timestamp"`
	Reserved1       string `json:"reserved1"`
	Reference       string `json:"reference"`
	Amount          int    `json:"amount"`
	ID              string `json:"id"`
	ForceClear      bool   `json:"force_clear"`
	Reserved4       string `json:"reserved4"`
	OrderDate       string `json:"order_date"`
	APIName         string `json:"api_name"`   // "general", "Redirect"
	Currency        string `json:"currency"`   // "THB"
	OrderType       string `json:"order_type"` //"Refund", "Sale"
	Reserved2       string `json:"reserved2"`
	Locked          bool   `json:"locked"`
	ChannelOrderID  string `json:"channel_order_id"`
	ErrorMessage    string `json:"error_message"`
	GatewayOrderID  string `json:"gateway_order_id"`
	Reserved3       string `json:"reserved3"`
	ErrorCode       string `json:"error_code"` // "SUCCESS", "FAIL", "DUPLICATED", "PENDING", "REFUNDED", "SIGNERROR"
	MID             string `json:"mid"`
	Acquirer        string `json:"acquirer"` // "Ksher"
	AcquirerOrderID string `json:"acquirer_order_id"`
	LogEntryURL     string `json:"log_entry_url"`
}
