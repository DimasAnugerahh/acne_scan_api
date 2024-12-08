package routes

import (
	"acne-scan-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (historyRoutes *HistoryRoutesImpl) HistoryRoutes(app *fiber.App) {
	app.Post("/history", middleware.UserMiddleware(), historyRoutes.HistoryHandlers.Create)
	app.Get("/history/:id", middleware.UserMiddleware(), middleware.UserMiddleware(), historyRoutes.HistoryHandlers.GetById)
	app.Get("/history", middleware.UserMiddleware(), historyRoutes.HistoryHandlers.GetAll)
}
