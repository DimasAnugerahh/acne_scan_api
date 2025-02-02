package validation

import (
	"acne-scan-api/internal/model/web"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func getValidationErrorMessage(field, tag, param string) string {
	switch tag {
		case "required":
			return fmt.Sprintf("%s is required", field)
		case "min":
			return fmt.Sprintf("%s must be at least %s", field, param)
		case "max":
			return fmt.Sprintf("%s must be at most %s", field, param)
		case "email":
			return fmt.Sprintf("%s must be a valid email address", field)
		case "url":
			return fmt.Sprintf("%s must be a valid URL", field)
		case "alpha":
			return fmt.Sprintf("%s must contain only alphabetical characters", field)
		case "alphanum":
			return fmt.Sprintf("%s must contain only alphanumeric characters", field)
		case "numeric":
			return fmt.Sprintf("%s must be a valid numeric value", field)
		case "len":
			return fmt.Sprintf("%s must be exactly %s characters long", field, param)
		// Add more cases for other validation tags as needed
		default:
			return fmt.Sprintf("validation error on field %s, tag %s", field, tag)
	}
}

func ValidationError(ctx *fiber.Ctx, err error) error {
	validationError, ok := err.(validator.ValidationErrors)
	if ok {
		errors := make(map[string][]string)

		for _, e := range validationError {
			field := e.Field()
			tag := e.Tag()
			param := e.Param()
			message := getValidationErrorMessage(field, tag, param)

			if _, exists := errors[field]; !exists {
				errors[field] = make([]string, 0)
			}

			errors[field] = append(errors[field], message)
		}

		response := web.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "validation failed",
			Errors:  errors,
		}

		return ctx.JSON(response)
	}

	return nil
}
