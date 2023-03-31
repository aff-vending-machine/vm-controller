package entity

import (
	"errors"
	"time"
)

type Slot struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Code      string    `json:"code"`
	Stock     int       `json:"stock"`
	Capacity  int       `json:"capacity"`
	Product   *Product  `json:"product,omitempty" gorm:"embedded;embeddedPrefix:product_"` // belong to
	IsEnable  bool      `json:"is_enable"`
}

func (e Slot) TableName() string {
	return "slots"
}

func (e Slot) Validate() error {
	if e.Code == "" {
		return errors.New("code is required")
	}
	if e.Stock < 0 {
		return errors.New("stock should be greater than or equal to zero")
	}
	if e.Capacity < 0 {
		return errors.New("capacity should be greater than or equal to zero")
	}

	return nil
}
