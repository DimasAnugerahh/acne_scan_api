package response

import (
	"acne-scan-api/internal/model/web"

	"github.com/gofiber/fiber/v2"
)

func SuccessResponse(c *fiber.Ctx, code int, message string, data any) error {
	return c.JSON(web.SuccessResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
