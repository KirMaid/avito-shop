package entities

import (
	"time"
)

type Purchase struct {
	ID        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	UserID    int       `json:"user_id"`
	MerchID   int       `json:"merch_id"`
	CreatedAt time.Time `json:"created_at"`
}
