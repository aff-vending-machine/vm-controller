package lugentpay

import (
	"bytes"
	"encoding/json"
	"io"
)

type QRCodeGenerateRequest struct {
	BillerCode   string `json:"billerCode"`   // LugentPay Biller Code
	BillerID     string `json:"billerID"`     // PromptPay Biller ID
	Reference1   string `json:"ref1"`         // Reference Code (Running)
	Reference2   string `json:"ref2"`         // Reference Code (Fix)
	Amount       string `json:"amount"`       // Transaction Amount with "." symbol
	StoreID      string `json:"storeID"`      // LugentPay Merchant Store ID
	TerminalID   string `json:"terminalID"`   // LugentPay Merchant Terminal ID
	MerchantName string `json:"merchantName"` // Merchant Name
	AccessCode   string `json:"accessCode"`   // Merchant Access Code
}

func (r *QRCodeGenerateRequest) ToReader() io.Reader {
	b, _ := json.Marshal(r)
	return bytes.NewBuffer(b)
}

type QRCodeGenerateResponse struct {
	ResponseCode        string `json:"res_code"`      // Response Code
	ResponseDescription string `json:"res_desc"`      // Response Description
	TransactionID       string `json:"transactionId"` // Transaction ID
	QRCode              string `json:"qrCode"`        // QR Code Data
}

func ToQRCodeGenerateResponse(decoder *json.Decoder) *QRCodeGenerateResponse {
	var res QRCodeGenerateResponse
	decoder.Decode(&res)
	return &res
}
