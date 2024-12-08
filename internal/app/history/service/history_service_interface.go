package service

import (
	"acne-scan-api/internal/app/history/repository"
	"acne-scan-api/internal/model/domain"
	"acne-scan-api/internal/model/web"

	"github.com/go-playground/validator"
)

type HistoryService interface {
	Create(request *web.HistoryRequest, historyJson []byte,userId int) error
	GetById(id int) (*domain.History, error)
	GetAll(id int) ([]*domain.History, error)
}

type HistoryServiceImpl struct {
	HistoryRepository repository.HistoryRepository
	Validator         *validator.Validate
}

func NewHistoryService(HistoryRepository repository.HistoryRepository, validate *validator.Validate) HistoryService {
	return &HistoryServiceImpl{HistoryRepository: HistoryRepository, Validator: validate}
}
