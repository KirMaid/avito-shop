package entities

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Balance  int    `json:"balance"`
	//Token          string    `json:"token"`
	//TokenExpiresAt time.Time `json:"token_expires_at"`
}
