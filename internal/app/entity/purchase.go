package entity

import (
	"time"
)

type Purchase struct {
	ID        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	UserID    int       `json:"user_id"`
	MerchID   int       `json:"merch_id"`
	//Quantity  int       `json:"quantity"`
	Total_price int       `json:"total_price"`
	CreatedAt   time.Time `json:"created_at"`
}
