package web

type ArticleResponse struct {
	ArticleId   int    `json:"article_id"`
	Image       string `json:"image"`
	Description string `json:"description"`
}