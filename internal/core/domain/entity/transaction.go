package entity

import (
	"time"
)

type Transaction struct {
	ID                  uint       `json:"id" gorm:"primarykey"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	MerchantOrderID     string     `json:"merchant_order_id" gorm:"not null,uniqueIndex"` // key to find order
	MachineSerialNumber string     `json:"machine_serial_number" gorm:"not null"`         // key to find machine
	RawCart             string     `json:"raw_cart"`                                      // ordered
	OrderQuantity       int        `json:"order_quantity" gorm:"not null"`                // ordered
	OrderPrice          float64    `json:"order_price" gorm:"not null"`                   // ordered
	OrderStatus         string     `json:"order_status" gorm:"not null"`                  // ordered
	OrderedAt           time.Time  `json:"ordered_at"`                                    // ordered
	PaymentChannel      string     `json:"payment"`                                       // ordered, key to find payment channel - MakeTransactionCreateRequest
	PaymentRequestedAt  *time.Time `json:"payment_requested_at"`                          // ordered - MakeTransactionCreateRequest
	Reference1          *string    `json:"reference1"`                                    // reference1 - MakeTransactionCreateResult
	Reference2          *string    `json:"reference2"`                                    // reference2
	Reference3          *string    `json:"reference3"`                                    // reference3
	CancelledBy         *string    `json:"cancelled_by"`                                  // cancelled - MakeTransactionCancel
	CancelledAt         *time.Time `json:"cancelled_at"`                                  // cancelled - MakeTransactionCancel
	ConfirmedPaidBy     *string    `json:"confirmed_paid_by" gorm:"default:null"`         // paid - MakeTransactionPaid
	ConfirmedPaidAt     *time.Time `json:"confirmed_paid_at" gorm:"default:null"`         // paid - MakeTransactionPaid
	RefError            *string    `json:"ref_error"`                                     // MakeTransactionError
	RefundAt            *time.Time `json:"refund_at" gorm:"default:null"`                 // refund
	RefundPrice         float64    `json:"refund_price"`                                  // refund
	ReceivedItemAt      *time.Time `json:"received_item_at" gorm:"default:null"`          // received - MakeTransactionDone
	ReceivedQuantity    int        `json:"received_quantity"`                             // received, refund - MakeTransactionDone
	PaidPrice           float64    `json:"paid_price"`                                    // received, refund - MakeTransactionDone
	Error               *string    `json:"error" gorm:"default:null"`                     // error - MakeTransactionError
	ErrorAt             *time.Time `json:"error_at" gorm:"default:null"`                  // MakeTransactionRefund
}

func (e *Transaction) TableName() string {
	return "transactions"
}
