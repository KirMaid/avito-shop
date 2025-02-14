package entities

type Inventory struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}
