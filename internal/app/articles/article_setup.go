package articles

import (
	articleHandlersPkg "acne-scan-api/internal/app/articles/handlers"
	articleRepositoryPkg "acne-scan-api/internal/app/articles/repository"
	articleRoutesPkg "acne-scan-api/internal/app/articles/routes"
	articleServicePkg "acne-scan-api/internal/app/articles/service"
	cloudstorage "acne-scan-api/internal/pkg/cloud_storage"
	"database/sql"

	"github.com/go-playground/validator"
)

func ArticleSetup(db *sql.DB, validate *validator.Validate,bs cloudstorage.StorageBucketUploader) articleRoutesPkg.ArticleRoutes{
	articleRepository:=articleRepositoryPkg.NewArticleRepository(db)
	articleService :=articleServicePkg.NewArticleService(articleRepository,validate,bs)
	articleHandlers:=articleHandlersPkg.NewArticleHandlerImpl(articleService)
	articleRoutes:=articleRoutesPkg.NewArticleRoutes(articleHandlers)

	return articleRoutes
}