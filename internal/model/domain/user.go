package domain

import "time"

type Users struct {
	User_id  uint   `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role" form:"role"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
