package smartedc

type SaleRequest struct {
	TradeType       string  `json:"trade_type" xml:"trade_type"`             // CARD
	Amount          float64 `json:"amount" xml:"amount"`                     // 100.25
	POSRefNo        string  `json:"pos_ref_no" xml:"pos_ref_no"`             // IV2018100100001
	TransactionType string  `json:"transaction_type" xml:"transaction_type"` // SALE
}

type SaleResult struct {
	POSRefNo         string  `json:"pos_ref_no" xml:"pos_ref_no"`                 // IV2018100100001
	ResponseCode     string  `json:"response_code" xml:"response_code"`           // 01
	ResponseMsg      string  `json:"response_msg" xml:"response_msg"`             // SUCCESS
	InvoiceNo        string  `json:"invoice_no" xml:"invoice_no"`                 // 000001
	CardNo           string  `json:"card_no" xml:"card_no"`                       // 444433XXXXXX9887
	Amount           float64 `json:"amount" xml:"amount"`                         // 100.25
	CardApprovalCode string  `json:"card_approval_code" xml:"card_approval_code"` // 123456
}
