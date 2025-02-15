package entities

type Inventory struct {
	UserID   int    `json:"user_id"`
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}
