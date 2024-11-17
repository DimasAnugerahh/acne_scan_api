package web

import "time"

type ArticleCreateRequest struct {
	ArticleId   int    `json:"article_id"`
	Image       string `json:"image" validate:"required"`
	Description string `json:"description" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ArticleUpdateRequest struct {
	Image       string `json:"image"`
	Description string `json:"description"`
}