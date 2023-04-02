package link2500

type RefundRequest struct {
	RefundPrice float64
}

type RefundResult struct {
	Payload                      []byte
	Result                       []byte
	ResponseText                 string
	MerchantName                 string
	TransactionDate              string
	TransactionTime              string
	ApprovalCode                 string
	InvoiceNumber                string
	TerminalIdentificationNumber string
	ISO8583MerchantNumber        string
	CardIssuerName               string
	PrimaryAccountNumber         string
	ExpirationDate               string
	BatchNumber                  string
	RetrievalReferenceNumber     string
	CardIssuerID                 string
	CardHolderName               string
	Amount                       string
}
