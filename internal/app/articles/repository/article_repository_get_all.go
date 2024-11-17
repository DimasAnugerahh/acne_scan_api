package repository

import (
	"acne-scan-api/internal/model/domain"
	"fmt"
)

func (articleRepository *ArticleRepositoryImpl) GetAll() ([]domain.Article, error) {

	result := []domain.Article{}

	// rows, err := articleRepository.DB.Query("select article_id, image, description, created_at,updated_at from article")

	rows, err := articleRepository.DB.Query("select article_id, image, description,created_at,updated_at from article")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		each := domain.Article{}
		err = rows.Scan(
			&each.ArticleId,
			&each.Image,
			&each.Description,
			&each.CreatedAt,
			&each.UpdatedAt,
		)
		if err != nil {
            fmt.Println(err.Error())

			return nil,err
        }
		result = append(result, each)
	}

	return result, nil
}
