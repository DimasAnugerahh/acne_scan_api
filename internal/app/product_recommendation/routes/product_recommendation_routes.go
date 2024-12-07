package routes

import (
	"acne-scan-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (ar *ProductRecommendationRoutesImpl) ProductRecommendationRoutes(app *fiber.App){
	app.Post("/recommendations",ar.productRecommendationHandler.Create)
	app.Get("/recommendations",middleware.UserMiddleware(),ar.productRecommendationHandler.GetAll)
	app.Get("/recommendations/:id",middleware.UserMiddleware(),ar.productRecommendationHandler.GetById)
	app.Patch("/recommendations/:id",ar.productRecommendationHandler.Update)
	app.Delete("/recommendations/:id",ar.productRecommendationHandler.Delete)



}
