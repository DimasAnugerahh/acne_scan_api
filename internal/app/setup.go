package app

import (
	"acne-scan-api/internal/app/articles"
	productrecommendation "acne-scan-api/internal/app/product_recommendation"
	cloudstorage "acne-scan-api/internal/pkg/cloud_storage"
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func InitApp(db *sql.DB, validate *validator.Validate, app *fiber.App, bs cloudstorage.StorageBucketUploader) {
	ArticleRoutes := articles.ArticleSetup(db, validate, bs)
	ProductRecommendationRoutes := productrecommendation.ProductRecommendationSetup(db, validate,bs)

	ProductRecommendationRoutes.ProductRecommendationRoutes(app)
	ArticleRoutes.ArticleRoutes(app)
}
