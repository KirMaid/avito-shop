package entities

type Good struct {
	ID    int    `json:"id" redis:"id"`
	Name  string `json:"name" redis:"name"`
	Price int    `json:"price" redis:"price"`
}
