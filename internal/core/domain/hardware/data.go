package hardware

import "encoding/json"

type Data struct {
	Mail             string
	MerchantOrderID  string
	InvoiceNo        string
	CardApprovalCode string
	Timestamp        string
	Cart             []Item
}

func (f *Data) TotalQuantity() int {
	total := 0
	for _, item := range f.Cart {
		total += item.Quantity
	}
	return total
}

func (f *Data) TotalPrice() float64 {
	price := 0.0
	for _, item := range f.Cart {
		price += float64(item.Quantity) * item.Price
	}
	return price
}

func (f *Data) TotalReceived() int {
	total := 0
	for _, item := range f.Cart {
		total += item.Received
	}
	return total
}

func (f *Data) TotalPay() float64 {
	price := 0.0
	for _, item := range f.Cart {
		price += float64(item.Received) * item.Price
	}
	return price
}

func (f *Data) Raw() string {
	b, _ := json.Marshal(f.Cart)
	return string(b)
}
