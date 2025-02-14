package entities

type UserInfo struct {
	Coins       int                 `json:"coins"`
	Inventory   []InventoryResponse `json:"inventory"`
	CoinHistory []CoinResponse      `json:"coinHistory"`
}

type InventoryResponse struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type CoinResponse struct {
	Received []ReceivedResponse `json:"received"`
	Sent     []SentResponse     `json:"sent"`
}

type ReceivedResponse struct {
	FromUser string `json:"fromUser"`
	Amount   int    `json:"amount"`
}

type SentResponse struct {
	ToUser string `json:"toUser"`
	Amount int    `json:"amount"`
}
