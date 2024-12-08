package web

import "time"

type HistoryResponse struct {
	HistoryId      int    `json:"history_id"`
	User_id        uint   `json:"user_id"`
	Name           string `json:"name"`
	Image          string `json:"image" form:"image"`
	Prediction     string `json:"prediction"`
	Recommendation string `json:"recommendation"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
