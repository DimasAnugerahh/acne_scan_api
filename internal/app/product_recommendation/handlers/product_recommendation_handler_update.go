package handlers

import (
	"acne-scan-api/internal/model/web"
	"acne-scan-api/internal/pkg/response"
	"acne-scan-api/internal/pkg/validation"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (pr *ProductRecommendationHandlerImpl) Update(c *fiber.Ctx) error {
	idparam := c.Params("id")
	id, err := strconv.Atoi(idparam)
	if err != nil {
		return response.BadRequest(c, "invalid product recommendation id", err)
	}

	ifExist, err := pr.productRecommendationService.GetById(id)
	if ifExist == nil {
		return response.BadRequest(c, "cannot found article to update", err)
	}

	req := new(web.ProductRecommendationUpdateRequest)
	if err := c.BodyParser(req); err != nil {
		return response.BadRequest(c, "failed to bind product recommendation request", err)
	}

	err = pr.productRecommendationService.Update(req, id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(c, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return response.BadRequest(c, "cannot found product recommendation to update", err)
		}
		return response.InternalServerError(c, "failed to update product recommendation, something happen", err.Error())
	}

	return response.StatusOk(c, 200, "product recommendation updated", nil)
}
