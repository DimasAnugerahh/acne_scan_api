package service

import (
	"acne-scan-api/internal/model/web"
	conversion "acne-scan-api/internal/pkg/conversion/request"
	"fmt"
)

func (articleService *ArticleServiceImpl) Create(request web.ArticleCreateRequest) error{
	var err error

	err = articleService.Validator.Struct(request)
	if err != nil {
		return err
	}

	article:=conversion.ArticleCreateRequestToArticleModel(request)

	err = articleService.ArticleRepository.Create(article)
		if err != nil {
		return fmt.Errorf("error when creating article %s", err.Error())
	}

	return nil
}