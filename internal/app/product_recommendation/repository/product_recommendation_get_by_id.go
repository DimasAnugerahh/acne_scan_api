package repository

import (
	"acne-scan-api/internal/model/domain"
	"fmt"
)

func (pr *ProductRecommendationImpl) GetById(id int) (*domain.ProductRecommendation,error) {

	result:=domain.ProductRecommendation{}

	err := pr.DB.QueryRow("select recommendation_id,image,link,description,created_at,updated_at from pruduct_recommendation where recommendation_id=?",id).Scan(
		&result.RecommendationId,
		&result.Image,
		&result.Link,
		&result.Description,
		&result.CreatedAt,
		&result.UpdatedAt,
		)
	if err != nil {
		return nil,fmt.Errorf("query gagal %s", err.Error())
	}

	return &result,nil
}
