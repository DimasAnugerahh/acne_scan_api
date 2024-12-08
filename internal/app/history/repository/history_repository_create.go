package repository

import (
	"acne-scan-api/internal/model/domain"
)

func (history *HistoryRepositoryImpl) Create(domainHistory *domain.History,historyJson []byte) error {
	stmt, err := history.DB.Prepare("insert into history (history_id,user_id,image,prediction,recommendation,created_at,updated_at) values (?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		domainHistory.HistoryId,
		domainHistory.User_id,
		string(historyJson),
		domainHistory.Prediction,
		domainHistory.Recommendation,
		domainHistory.CreatedAt,
		domainHistory.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
