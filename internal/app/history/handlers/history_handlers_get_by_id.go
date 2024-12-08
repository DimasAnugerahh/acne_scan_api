package handlers

import (
	"acne-scan-api/internal/pkg/response"
	"acne-scan-api/internal/pkg/validation"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (history *HistoryHandlersImpl) GetById(c *fiber.Ctx) error {
	idparam := c.Params("id")
	id, err := strconv.Atoi(idparam)
	if err != nil {
		return response.BadRequest(c, "invalid history users id", err)
	}

	data, err := history.service.GetById(id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(c, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return response.BadRequest(c, "history not found", err)
		}
		return response.InternalServerError(c, "failed to get history, something happen", err.Error())
	}

	return response.StatusOk(c, http.StatusOK, "success", data)
}
