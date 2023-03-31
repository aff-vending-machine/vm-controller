package entity

import "time"

type Customer struct {
	ID         uint      `json:"id" gorm:"primarykey"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Email      string    `json:"email"`
	Tel        string    `json:"tel"`
	Cart       string    `json:"cart"`
	IsReceived bool      `json:"is_received"`
}

func (e *Customer) TableName() string {
	return "customers"
}
