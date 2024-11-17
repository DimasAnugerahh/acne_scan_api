package repository

import (
	"acne-scan-api/internal/model/domain"
	"database/sql"
)

type ArticleRepository interface {
	Create(article *domain.Article)error
	// Update(article *domain.Article,id int)error
	// Delete(article *domain.Article,id int)error
	GetAll()([]domain.Article,error)
	GetById(id int)(*domain.Article,error)
}

type ArticleRepositoryImpl struct{
	DB *sql.DB
}

func NewArticleRepository(db *sql.DB)ArticleRepository{
	return &ArticleRepositoryImpl{DB: db}
}