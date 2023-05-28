package link2500

type Result struct {
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
	MerchantNumber               string
	BatchNumber                  string
	RetrievalReferenceNumber     string
	CardIssuerID                 string
	CardHolderName               string
	Amount                       string
	Hosts                        []string
}

func (r *Result) ToSaleResult(payload []byte, result []byte) *SaleResult {
	return &SaleResult{
		Payload:                      payload,
		Result:                       result,
		ResponseText:                 r.ResponseText,
		MerchantName:                 r.MerchantName,
		TransactionDate:              r.TransactionDate,
		TransactionTime:              r.TransactionTime,
		ApprovalCode:                 r.ApprovalCode,
		InvoiceNumber:                r.InvoiceNumber,
		TerminalIdentificationNumber: r.TerminalIdentificationNumber,
		ISO8583MerchantNumber:        r.ISO8583MerchantNumber,
		CardIssuerName:               r.CardIssuerName,
		PrimaryAccountNumber:         r.PrimaryAccountNumber,
		ExpirationDate:               r.ExpirationDate,
		BatchNumber:                  r.BatchNumber,
		RetrievalReferenceNumber:     r.RetrievalReferenceNumber,
		CardIssuerID:                 r.CardIssuerID,
		CardHolderName:               r.CardHolderName,
		Amount:                       r.Amount,
	}
}

func (r *Result) ToVoidResult(payload []byte, result []byte) *VoidResult {
	return &VoidResult{
		Payload:                      payload,
		Result:                       result,
		ResponseText:                 r.ResponseText,
		MerchantName:                 r.MerchantName,
		TransactionDate:              r.TransactionDate,
		TransactionTime:              r.TransactionTime,
		ApprovalCode:                 r.ApprovalCode,
		InvoiceNumber:                r.InvoiceNumber,
		TerminalIdentificationNumber: r.TerminalIdentificationNumber,
		ISO8583MerchantNumber:        r.ISO8583MerchantNumber,
		CardIssuerName:               r.CardIssuerName,
		PrimaryAccountNumber:         r.PrimaryAccountNumber,
		ExpirationDate:               r.ExpirationDate,
		BatchNumber:                  r.BatchNumber,
		RetrievalReferenceNumber:     r.RetrievalReferenceNumber,
		CardIssuerID:                 r.CardIssuerID,
		CardHolderName:               r.CardHolderName,
		Amount:                       r.Amount,
	}
}

func (r *Result) ToRefundResult(payload []byte, result []byte) *RefundResult {
	return &RefundResult{
		Payload:                      payload,
		Result:                       result,
		ResponseText:                 r.ResponseText,
		MerchantName:                 r.MerchantName,
		TransactionDate:              r.TransactionDate,
		TransactionTime:              r.TransactionTime,
		ApprovalCode:                 r.ApprovalCode,
		InvoiceNumber:                r.InvoiceNumber,
		TerminalIdentificationNumber: r.TerminalIdentificationNumber,
		ISO8583MerchantNumber:        r.ISO8583MerchantNumber,
		CardIssuerName:               r.CardIssuerName,
		PrimaryAccountNumber:         r.PrimaryAccountNumber,
		ExpirationDate:               r.ExpirationDate,
		BatchNumber:                  r.BatchNumber,
		RetrievalReferenceNumber:     r.RetrievalReferenceNumber,
		CardIssuerID:                 r.CardIssuerID,
		CardHolderName:               r.CardHolderName,
		Amount:                       r.Amount,
	}
}

func (r *Result) ToSettlementResult(payload []byte, result []byte) *SettlementResult {
	return &SettlementResult{
		Payload: payload,
		Result:  result,
		Hosts:   r.Hosts,
	}
}
