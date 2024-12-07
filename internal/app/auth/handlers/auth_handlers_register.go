package handlers

import (
	"acne-scan-api/internal/model/web"
	"acne-scan-api/internal/pkg/response"
	"acne-scan-api/internal/pkg/validation"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (auth *AuthHandlersImpl) Register(c *fiber.Ctx) error {
	req := new(web.Register)

	if err := c.BodyParser(req); err != nil {
		fmt.Println(err.Error())
		return response.BadRequest(c, "failed to bind register", err)
	}

	err:=auth.service.Register(req)
	if err!=nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(c, err)
		}
		return response.InternalServerError(c, "failed to register, something happen", err.Error())
	}

	return response.StatusCreated(c, http.StatusCreated, "success to register", nil)
}
