package routes

import "github.com/gofiber/fiber/v2"

func (historyRoutes *HistoryRoutesImpl) HistoryRoutes(app *fiber.App) {
	app.Post("/history", historyRoutes.HistoryHandlers.Create)
}
