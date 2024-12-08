package routes

import (
	"acne-scan-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (ar *ProductRecommendationRoutesImpl) ProductRecommendationRoutes(app *fiber.App) {
	app.Post("/recommendations", middleware.AdminMiddleware(), ar.productRecommendationHandler.Create)
	app.Get("/recommendations", ar.productRecommendationHandler.GetAll)
	app.Get("/recommendations/:id", ar.productRecommendationHandler.GetById)
	app.Patch("/recommendations/:id", middleware.AdminMiddleware(), ar.productRecommendationHandler.Update)
	app.Delete("/recommendations/:id", middleware.AdminMiddleware(), ar.productRecommendationHandler.Delete)

}
