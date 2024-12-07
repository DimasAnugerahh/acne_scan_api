package routes

import (
	"acne-scan-api/internal/app/auth/handlers"

	"github.com/gofiber/fiber/v2"
)

type AuthRoutes interface {
	AuthRoutes(app *fiber.App)
}

type AuthRoutesImpl struct{
	authHandlers handlers.AuthHandlers
}

func NewAuthRoutes (authHandlres handlers.AuthHandlers)AuthRoutes{
	return &AuthRoutesImpl{authHandlers: authHandlres}
}