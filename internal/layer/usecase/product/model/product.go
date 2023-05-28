package model

type Product struct {
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	Group    string  `json:"group"`
	ImageURL string  `json:"image_url"`
	Price    float64 `json:"price"`
}
