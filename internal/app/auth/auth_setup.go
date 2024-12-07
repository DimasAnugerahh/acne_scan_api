package auth

import (
	"acne-scan-api/internal/app/auth/handlers"
	"acne-scan-api/internal/app/auth/repository"
	RoutesPKG "acne-scan-api/internal/app/auth/routes"
	"acne-scan-api/internal/app/auth/service"
	"database/sql"

	"github.com/go-playground/validator"
)

func AuthSetup(db *sql.DB,validate *validator.Validate)RoutesPKG.AuthRoutes{
	authRepository:=repository.NewAuthRepository(db)
	authService:=service.NewAuthService(authRepository,validate)
	authHandlers:=handlers.NewAuthHandlers(authService)
	authRoutes:=RoutesPKG.NewAuthRoutes(authHandlers)

	return authRoutes
}