package handlers

import (
	"acne-scan-api/internal/model/web"
	"acne-scan-api/internal/pkg/response"
	"acne-scan-api/internal/pkg/validation"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (articleHandler *ArticleHandlerImpl) Create(c *fiber.Ctx) error {
	req := new(web.ArticleCreateRequest)
	
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "failed to bind category request", err)
	}

	err := articleHandler.ArticleService.Create(*req)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(c, err)
		}
		return response.InternalServerError(c, "failed to create category, something happen", err)
	}

	return response.StatusCreated(c, http.StatusCreated, "success to create category", nil)
}
