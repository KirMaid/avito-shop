package entities

type Inventory struct {
	UserID   int `json:"user_id"`
	GoodID   int `json:"good_id"`
	Quantity int `json:"quantity"`
}
