package hardware

type Item struct {
	SlotCode string  `json:"code"`
	Name     string  `json:"name"`
	ImageURL string  `json:"image_url"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Received int     `json:"received"`
}

func (item *Item) Clear() {
	item.SlotCode = ""
	item.Name = ""
	item.ImageURL = ""
	item.Quantity = 0
	item.Price = 0
	item.Received = 0
}
