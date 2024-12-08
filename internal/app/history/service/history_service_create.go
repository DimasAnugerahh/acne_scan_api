package service

import (
	"acne-scan-api/internal/model/web"
	conversion "acne-scan-api/internal/pkg/conversion/request"
	"fmt"
	"time"
)

func (history *HistoryServiceImpl) Create(request *web.HistoryRequest,historyJson []byte) error {
	var err error

	err = history.Validator.Struct(request)
	if err != nil {
		return err
	}

	historyConv := conversion.HistoryRequestToModel(request)

	wib, err := time.LoadLocation("Asia/Jakarta") // WIB (UTC+7)
	if err != nil {
		return fmt.Errorf("error loading WIB location: %s", err.Error())
	}

	createdAt := time.Now().In(wib)
	historyConv.CreatedAt = createdAt
	historyConv.UpdatedAt = createdAt

	err = history.HistoryRepository.Create(historyConv,historyJson)
	if err != nil {
		return fmt.Errorf("error register %s", err.Error())
	}

	return nil
}
