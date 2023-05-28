package entity

import (
	"errors"
)

type Product struct {
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	Group    string  `json:"group"`
	ImageURL string  `json:"image_url"`
	Price    float64 `json:"price"`
}

func (e Product) TableName() string {
	return "products"
}

func (e Product) Validate() error {
	if e.Name == "" {
		return errors.New("name is required")
	}
	if e.SKU == "" {
		return errors.New("sku is required")
	}
	if e.Price < 0 {
		return errors.New("price should be greater than or equal to zero")
	}
	return nil
}
