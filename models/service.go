package models

import "time"

type Service struct {
	ServiceID     string    `json:"service_id"`
	Name          string    `json:"name"`
	Price         float64   `json:"price"`
	Category      string    `json:"category"`
	PickupAddress string    `json:"pickup_address"`
	DestAddress   string    `json:"dest_address"`
	CreateAt      time.Time `json:"created_at"`
}
