package routes

import (
	"acne-scan-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (articleRoutes *ArticleRoutesImpl) ArticleRoutes(app *fiber.App) {
	app.Post("/articles", articleRoutes.ArticleHandler.Create)
	app.Get("/articles",middleware.UserMiddleware(), articleRoutes.ArticleHandler.GetAll)
	app.Get("/articles/:id",middleware.UserMiddleware(), articleRoutes.ArticleHandler.GetById)
	app.Delete("/articles/:id", articleRoutes.ArticleHandler.Delete)
	app.Patch("/articles/:id", articleRoutes.ArticleHandler.Update)

}