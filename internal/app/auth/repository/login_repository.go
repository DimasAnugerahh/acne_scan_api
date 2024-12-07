package repository

import (
	"acne-scan-api/internal/model/domain"
)

func (login *AuthRepositoryImpl) Login(username, password string) (*domain.Users, error) {
	result := domain.Users{}

	err:=login.DB.QueryRow("select username,password,role from users where username=? and password=?",username,password).Scan(
		&result.Username,
		&result.Password,
		&result.Role,

	)
	if err != nil {
		return nil, err
	}

	return &result,nil
}
