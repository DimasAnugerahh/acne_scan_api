package routes

import (
	"acne-scan-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (articleRoutes *ArticleRoutesImpl) ArticleRoutes(app *fiber.App) {
	app.Post("/articles",middleware.AdminMiddleware(), articleRoutes.ArticleHandler.Create)
	app.Get("/articles", articleRoutes.ArticleHandler.GetAll)
	app.Get("/articles/:id", articleRoutes.ArticleHandler.GetById)
	app.Delete("/articles/:id",middleware.AdminMiddleware(), articleRoutes.ArticleHandler.Delete)
	app.Patch("/articles/:id",middleware.AdminMiddleware(), articleRoutes.ArticleHandler.Update)

}