package repository

import (
	"acne-scan-api/internal/model/domain"
	"database/sql"
)

type AuthRepository interface {
	Login(username,password string)(*domain.Users,error)
	Register(users *domain.Users) error
}

type AuthRepositoryImpl struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB)AuthRepository{
	return &AuthRepositoryImpl{DB:db}
}