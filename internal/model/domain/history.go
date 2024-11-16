package domain

import "time"

type History struct {
	HistoryId int       `gorm:"type:int;primarykey" json:"history_id"`
	Image     string    `gorm:"type:varchar(255)" json:"image"`
	Result    string    `gorm:"type:varchar(255)" json:"result"`
	User_id   uint      `gorm:"type:int" json:"user_id"`
	
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
