package handlers

import (
	"acne-scan-api/internal/model/web"
	"acne-scan-api/internal/pkg/response"
	"acne-scan-api/internal/pkg/validation"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (history *HistoryHandlersImpl) Create(c *fiber.Ctx) error {
	req := new(web.HistoryRequest)

	if err := c.BodyParser(req); err != nil {
		fmt.Println(err.Error())
		return response.BadRequest(c, "failed to bind history request", err)
	}

	historyJson, err := json.Marshal(req.Image)
	if err != nil {
		log.Fatalf("Error serializing map to JSON: %v", err)
	}

	err = history.service.Create(req,historyJson)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(c, err)
		}
		return response.InternalServerError(c, "failed to create history, something happen", err.Error())
	}

	return response.StatusCreated(c, http.StatusCreated, "success to create history", nil)
}
