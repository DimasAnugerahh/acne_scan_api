package handlers

import (
	"acne-scan-api/internal/app/auth/service"

	"github.com/gofiber/fiber/v2"
)

type AuthHandlers interface {
	Login(ctx *fiber.Ctx)  error
	Register(c *fiber.Ctx) error
}

type AuthHandlersImpl struct {
	service service.AuthService
}

func NewAuthHandlers(authService service.AuthService) AuthHandlers {
	return &AuthHandlersImpl{service: authService}
}