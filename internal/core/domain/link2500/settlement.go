package link2500

type SettlementRequest struct {
	MerchantID string `json:"merchant_id"`
}

type SettlementResult struct {
	Payload []byte
	Result  []byte
	Hosts   []string
}
