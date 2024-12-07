package service

import (
	"acne-scan-api/internal/model/domain"
)

func (authService *AuthServiceImpl) Login(username, password string) (*domain.Users, error) {
	data, err := authService.AuthRepository.Login(username, password)
	if err != nil || data == nil {
		return nil, err
	}

	return data, nil
}
