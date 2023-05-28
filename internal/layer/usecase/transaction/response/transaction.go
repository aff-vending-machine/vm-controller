package response

import (
	"time"

	"github.com/aff-vending-machine/vm-controller/internal/core/domain/entity"
)

type Transaction struct {
	ID                  uint       `json:"id"`
	MerchantOrderID     string     `json:"merchant_order_id"`     // key to find order
	MachineSerialNumber string     `json:"machine_serial_number"` // key to find machine
	Location            string     `json:"location"`              // ordered
	RawCart             string     `json:"raw_cart"`              // ordered
	OrderQuantity       int        `json:"order_quantity"`        // ordered
	OrderPrice          float64    `json:"order_price"`           // ordered
	OrderStatus         string     `json:"order_status"`          // ordered
	OrderedAt           time.Time  `json:"ordered_at"`            // ordered
	PaymentChannel      string     `json:"payment_channel"`       // ordered, key to find payment channel - MakeTransactionCreateRequest
	PaymentRequestedAt  *time.Time `json:"payment_requested_at"`  // ordered - MakeTransactionCreateRequest
	Reference1          *string    `json:"reference1"`            // reference1 - MakeTransactionCreateResult
	Reference2          *string    `json:"reference2"`            // reference2
	Reference3          *string    `json:"reference3"`            // reference3
	CancelledBy         *string    `json:"cancelled_by"`          // cancelled - MakeTransactionCancel
	CancelledAt         *time.Time `json:"cancelled_at"`          // cancelled - MakeTransactionCancel
	ConfirmedPaidBy     *string    `json:"confirmed_paid_by"`     // paid - MakeTransactionPaid
	ConfirmedPaidAt     *time.Time `json:"confirmed_paid_at"`     // paid - MakeTransactionPaid
	RefError            *string    `json:"ref_error"`             // MakeTransactionError
	RefundAt            *time.Time `json:"refund_at"`             // refund
	RefundPrice         float64    `json:"refund_price"`          // refund
	ReceivedItemAt      *time.Time `json:"received_item_at"`      // received - MakeTransactionDone
	ReceivedQuantity    int        `json:"received_quantity"`     // received, refund - MakeTransactionDone
	PaidPrice           float64    `json:"paid_price"`            // received, refund - MakeTransactionDone
	IsError             bool       `json:"is_error"`              // error
	Error               *string    `json:"error"`                 // error - MakeTransactionError
	ErrorAt             *time.Time `json:"error_at"`              // MakeTransactionRefund
}

func ToTransaction(e *entity.Transaction) *Transaction {
	return &Transaction{
		ID:                  e.ID,
		MerchantOrderID:     e.MerchantOrderID,
		MachineSerialNumber: e.MachineSerialNumber,
		Location:            e.Location,
		RawCart:             e.RawCart,
		OrderQuantity:       e.OrderQuantity,
		OrderPrice:          e.OrderPrice,
		OrderStatus:         e.OrderStatus,
		OrderedAt:           e.OrderedAt,
		PaymentChannel:      e.PaymentChannel,
		PaymentRequestedAt:  e.PaymentRequestedAt,
		Reference1:          e.Reference1,
		Reference2:          e.Reference2,
		Reference3:          e.Reference3,
		CancelledBy:         e.CancelledBy,
		CancelledAt:         e.CancelledAt,
		ConfirmedPaidBy:     e.ConfirmedPaidBy,
		ConfirmedPaidAt:     e.ConfirmedPaidAt,
		RefError:            e.RefError,
		RefundAt:            e.RefundAt,
		RefundPrice:         e.RefundPrice,
		ReceivedItemAt:      e.ReceivedItemAt,
		ReceivedQuantity:    e.ReceivedQuantity,
		PaidPrice:           e.PaidPrice,
		IsError:             e.IsError,
		Error:               e.Error,
		ErrorAt:             e.ErrorAt,
	}
}

func ToTransactionList(entities []entity.Transaction) []Transaction {
	results := make([]Transaction, len(entities))
	for i, e := range entities {
		results[i] = *ToTransaction(&e)
	}

	return results
}
