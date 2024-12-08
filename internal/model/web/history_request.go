package web

type Images map[string]string

type HistoryRequest struct {
	User_id        uint   `json:"user_id"`
	Image          Images `json:"image" form:"image"`
	Prediction     string `json:"prediction"`
	Recommendation string `json:"recommendation"`
}
