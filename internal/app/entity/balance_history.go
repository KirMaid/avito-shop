package entity

import "time"

type BalanceHistory struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	Change_amount  int       `json:"change_amount"`
	Operation_type string    `json:"operation_type"`
	Created_at     time.Time `json:"created_at"`
}
