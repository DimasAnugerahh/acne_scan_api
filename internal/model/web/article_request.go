package web

type ArticleCreateRequest struct {
	ArticleId   int    `json:"article_id"`
	Image       string `json:"image" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type ArticleUpdateRequest struct {
	Image       string `json:"image"`
	Description string `json:"description"`
}