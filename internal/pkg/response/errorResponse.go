package response

import (
	"acne-scan-api/internal/model/web"

	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(c *fiber.Ctx, code int, message string,errors any) error{

	return c.JSON(web.ErrorResponse{
		Code: code,
		Message: message,
		Errors: errors,
	})

}