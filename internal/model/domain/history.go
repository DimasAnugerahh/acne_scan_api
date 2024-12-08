package domain

import "time"

type Images map[string]string

type History struct {
	HistoryId      int    `json:"history_id"`
	User_id        uint   `json:"user_id"`
	Image          Images `json:"image" form:"image"`
	Prediction     string `json:"prediction"`
	Recommendation string `json:"recommendation"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
