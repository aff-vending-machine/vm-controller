package entity

import "time"

type Machine struct {
	ID           uint      `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `json:"name"`
	SerialNumber string    `json:"serial_number" gorm:"uniqueIndex"`
	Branch       string    `json:"branch"`
	Location     string    `json:"location"`
	Type         string    `json:"type"`
	Vendor       string    `json:"vendor"`
	Center       string    `json:"center"`
	Status       string    `json:"status"`
	TestMode     bool      `json:"test_mode"`
}

func (e *Machine) TableName() string {
	return "machines"
}
