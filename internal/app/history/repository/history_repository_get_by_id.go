package repository

import (
	"acne-scan-api/internal/model/domain"
	"encoding/json"
	"fmt"
)

func (history *HistoryRepositoryImpl) GetById(id int) (*domain.History, error) {
	result := domain.History{}

	var imageData []byte

	err := history.DB.QueryRow("select history.history_id, history.user_id,users.username, history.image, history.prediction, history.recommendation, history.created_at, history.updated_at from history left join users on users.user_id=history.user_id where history_id=?", id).Scan(
		&result.HistoryId,
		&result.User_id,
		&result.Name,
		&imageData,
		&result.Prediction,
		&result.Recommendation,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if err := json.Unmarshal(imageData, &result.Image); err != nil {
		return nil, fmt.Errorf("failed to parse image JSON: %v", err)
	}

	return &result, nil
}
