package app

import (
	"acne-scan-api/internal/app/articles"
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func InitApp(db *sql.DB,validate *validator.Validate,app *fiber.App){
	ArticleRoutes := articles.ArticleSetup(db,validate)

	ArticleRoutes.ArticleRoutes(app)
}