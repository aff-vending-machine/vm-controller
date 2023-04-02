package ksher

type CreateOrderBody struct {
	MerchantOrderID string `json:"merchant_order_id" validate:"required"` // The unique order number used by the merchant
	Amount          int    `json:"amount" validate:"required"`            // How much to charge.
	Note            string `json:"note,omitempty"`                        // Order note
	Timestamp       string `json:"timestamp" validate:"required"`         // Timestamp
	MID             string `json:"mid,omitempty"`                         // The merchant ID of this order is only valid when there are multiple sub-merchants under the same account.
	Provider        string `json:"provider,omitempty"`                    // Payment gateway, optional values ('th','ph','my')
	Channel         string `json:"channel" validate:"required"`           // Payment Merchant support.
	DeviceID        string `json:"device_id,omitempty"`                   // terminal id from which the request is sent, assigned by merchant.
	OperatoreID     string `json:"operator_id,omitempty"`                 // operator_id number at cashier, using for merchant have muti level account or Cashier. For more information please check at Muti level account or Cashier
	Signature       string `json:"signature" validate:"required"`         // Request signature
}

type CreateOrderResult struct {
	Acquirer        string `json:"acquirer"` // "Ksher"
	AcquirerOrderID string `json:"acquirer_order_id"`
	Amount          int    `json:"amount"`
	APIName         string `json:"api_name"` // "general", "Redirect", "CscanB"
	Channel         string `json:"channel"`  // wechat, truemoney
	Cleared         bool   `json:"cleared"`
	Currency        string `json:"currency"`   // "THB"
	ErrorCode       string `json:"error_code"` // "REFUNDED", "SUCCESS"
	ErrorMessage    string `json:"error_message"`
	ForceClear      bool   `json:"force_clear"`
	GatewayOrderID  string `json:"gateway_order_id"`
	ID              string `json:"id"`
	Locked          bool   `json:"locked"`
	MerchantOrderID string `json:"merchant_order_id"`
	Note            string `json:"note"`
	OrderType       string `json:"order_type"` //"Refund", "Sale", "Available"
	Reference       string `json:"reference"`
	Reserved1       string `json:"reserved1"`
	Signature       string `json:"signature"`
	Status          string `json:"status"` // "Refunded", "Closed"
	Timestamp       string `json:"timestamp"`
}
