package model

type Machine struct {
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
	Branch       string `json:"branch"`
	Location     string `json:"location"`
	Type         string `json:"type"`
	Vendor       string `json:"vendor"`
	Center       string `json:"center"`
	Status       string `json:"status"`
}
