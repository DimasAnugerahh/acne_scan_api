package domain

import "time"

type ProductRecommendation struct {
	RecommendationId int    `gorm:"type:int;primarykey" json:"recommendation_id"`
	Image            string `gorm:"type:varchar(255)" json:"image"`
	Link             string `gorm:"type:varchar(255)" json:"link"`
	Description      string `gorm:"type:varchar(512)" json:"description"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}