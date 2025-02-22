package usecases

type UserInfoDTO struct {
	Coins       int            `json:"coins"`
	Inventory   []InventoryDTO `json:"inventory"`
	CoinHistory CoinHistoryDTO `json:"coinHistory"`
}

type InventoryDTO struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type CoinHistoryDTO struct {
	Received []ReceivedDTO `json:"received"`
	Sent     []SentDTO     `json:"sent"`
}

type ReceivedDTO struct {
	FromUser string `json:"fromUser"`
	Amount   int    `json:"amount"`
}

type SentDTO struct {
	ToUser string `json:"toUser"`
	Amount int    `json:"amount"`
}
