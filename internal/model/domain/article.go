package domain

import "time"

type Article struct {
	ArticleId   int    `json:"article_id"`
	Image       string `json:"image"`
	Description string `json:"description"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}