package service

import (
	"acne-scan-api/internal/app/auth/repository"
	"acne-scan-api/internal/model/domain"

	"github.com/go-playground/validator"
)

type AuthService interface {
	Login(username, password string) (*domain.Users, error)
}

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	Validator   *validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, validate *validator.Validate)AuthService{
	return &AuthServiceImpl{AuthRepository: authRepository,Validator: validate}
}
