package smartedc

type VoidRequest struct {
	TradeType        string  `json:"trade_type" xml:"trade_type"`                 // CARD
	InvoiceNo        string  `json:"invoice_no" xml:"invoice_no"`                 // 000001
	CardApprovalCode string  `json:"card_approval_code" xml:"card_approval_code"` // 123456
	POSRefNo         string  `json:"pos_ref_no" xml:"pos_ref_no"`                 // IV2018100100001
	Amount           float64 `json:"amount" xml:"amount"`                         // 100.25
	TransactionType  string  `json:"transaction_type" xml:"transaction_type"`     // VOID
}

type VoidResult struct {
	POSRefNo      string  `json:"pos_ref_no" xml:"pos_ref_no"`         // IV2018100100001
	ResponseCode  string  `json:"response_code" xml:"response_code"`   // 01
	ResponseMsg   string  `json:"response_msg" xml:"response_msg"`     // SUCCESS
	TransactionID string  `json:"transaction_id" xml:"transaction_id"` // 1217752533368018012014070332
	InvoiceNo     string  `json:"invoice_no" xml:"invoice_no"`         // 000001
	Amount        float64 `json:"amount" xml:"amount"`                 // 100.25
}
