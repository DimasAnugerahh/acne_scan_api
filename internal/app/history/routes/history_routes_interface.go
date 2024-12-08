package routes

import (
	"acne-scan-api/internal/app/history/handlers"

	"github.com/gofiber/fiber/v2"
)

type HistoryRoutes interface {
	HistoryRoutes(app *fiber.App)
}

type HistoryRoutesImpl struct{
	HistoryHandlers handlers.HistoryHandlers
}

func NewHistoryRoutes (HistoryHandlres handlers.HistoryHandlers)HistoryRoutes{
	return &HistoryRoutesImpl{HistoryHandlers: HistoryHandlres}
}