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

// TODO Поменять int на string обратно как протестирую прототип
type ReceivedDTO struct {
	FromUser int `json:"fromUser"`
	Amount   int `json:"amount"`
}

type SentDTO struct {
	ToUser int `json:"toUser"`
	Amount int `json:"amount"`
}
