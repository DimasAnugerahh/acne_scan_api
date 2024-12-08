package handlers

import (
	"acne-scan-api/internal/app/history/service"

	"github.com/gofiber/fiber/v2"
)

type HistoryHandlers interface {
	Create(c *fiber.Ctx) error 
}

type HistoryHandlersImpl struct {
	service service.HistoryService
}

func NewHistoryHandlers(HistoryService service.HistoryService) HistoryHandlers {
	return &HistoryHandlersImpl{service: HistoryService}
}