package lugentpay

import (
	"bytes"
	"encoding/json"
	"io"
)

type InquiryBody struct {
	TransactionID string `json:"transactionId"` // Transaction ID from QR Generate Response
	AccessCode    string `json:"accessCode"`    // Merchant Access Code
}

func (r *InquiryBody) ToReader() io.Reader {
	b, _ := json.Marshal(r)
	return bytes.NewBuffer(b)
}

type InquiryResult struct {
	ResponseCode          string `json:"res_code"`      // Response Code
	ResponseDescription   string `json:"res_desc"`      // Response Description
	TransactionID         string `json:"transactionId"` // Transaction ID
	PaymentConfirmationID string `json:"paymentId"`     // QR Code Data
}

func ToInquiry(decoder *json.Decoder) *InquiryResult {
	var res InquiryResult
	decoder.Decode(&res)
	return &res
}
