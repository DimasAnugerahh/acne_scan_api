package web

type UserLogin struct {
	User_id  uint   `json:"user_id" form:"user_id"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Role     string `json:"role" form:"role"`
}