package domain

type Users struct {
	ID       uint   `gorm:"type:int;primarykey" json:"id"`
	Username string `gorm:"type:varchar(255)" json:"username"`
	Email    string `gorm:"type:varchar(255)" json:"email"`
}