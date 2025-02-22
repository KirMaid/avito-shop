package entities

type Inventory struct {
	UserID   int `json:"user_id" redis:"user_id"`
	GoodID   int `json:"good_id" redis:"good_id"`
	Quantity int `json:"quantity" redis:"quantity"`
}
