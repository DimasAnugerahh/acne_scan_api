package web

type AuthResponse struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
	UserId   int    `json:"user_id"`
}
