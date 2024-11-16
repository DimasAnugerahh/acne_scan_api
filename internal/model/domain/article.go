package domain

import "time"

type Article struct {
	ArticleId   int    `gorm:"type:int;primarykey" json:"article_id"`
	Image       string `gorm:"type:varchar(255)" json:"image"`
	Description string `gorm:"type:varchar(512)" json:"description"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}