package handlers

import (
	"acne-scan-api/internal/model/web"
	conversion "acne-scan-api/internal/pkg/conversion/response"
	"acne-scan-api/internal/pkg/jwt"
	"acne-scan-api/internal/pkg/response"
	"acne-scan-api/internal/pkg/validation"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (authHandlers *AuthHandlersImpl) Login(c *fiber.Ctx) error {
	req := new(web.UserLogin)
	if err := c.BodyParser(req); err != nil {

		return response.BadRequest(c, "failed to login request", err)
	}

	data, err := authHandlers.service.Login(req.Username, req.Password)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(c, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return response.BadRequest(c, "not found", err)
		}
		return response.InternalServerError(c, "failed to login, something happer", err.Error())
	}

	token, err := jwt.GenerateAccessToken(data)
	if err != nil {
		return fmt.Errorf("failed to generate token %s",err.Error())
	}


	result:=conversion.AuthResponse(*req, token,data.Role)

	return response.StatusOk(c, 200, "login berhasil", result)
}
