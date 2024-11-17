package routes

import "github.com/gofiber/fiber/v2"

func (articleRoutes *ArticleRoutesImpl) ArticleRoutes(app *fiber.App) {
	app.Post("/articles", articleRoutes.ArticleHandler.Create)
}