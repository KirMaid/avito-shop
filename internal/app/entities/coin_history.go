package entities

import "time"

type CoinHistory struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	ChangeAmount  int       `json:"change_amount"`
	OperationType string    `json:"operation_type"`
	CreatedAt     time.Time `json:"created_at"`
}
