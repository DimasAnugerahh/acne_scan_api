package conversion

import (
	"acne-scan-api/internal/model/domain"
	"acne-scan-api/internal/model/web"
)

func HistoryRequestToModel(request *web.HistoryRequest) *domain.History {
	return &domain.History{
		User_id:        request.User_id,
		Image:          domain.Images(request.Image),
		Prediction:     request.Prediction,
		Recommendation: request.Recommendation,
	}
}
