package entities

type User struct {
	ID       int    `json:"id" redis:"id"`
	Username string `json:"username" redis:"username"`
	Password string `json:"password" redis:"password"`
	Balance  int    `json:"balance" redis:"balance"`
}
