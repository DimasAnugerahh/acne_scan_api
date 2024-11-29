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

func (articleHandler *ArticleHandlerImpl) Create(c *fiber.Ctx) error {
	req := new(web.ArticleCreateRequest)
	
	if err := c.BodyParser(req); err != nil {
		fmt.Println(err.Error())
		return response.BadRequest(c, "failed to bind article request", err)
	}

	image,err:=c.FormFile("image")
	if err!=nil {
		return err
	}

	err = articleHandler.ArticleService.Create(*req,image,c)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(c, err)
		}
		return response.InternalServerError(c, "failed to create article, something happen", err.Error())
	}

	return response.StatusCreated(c, http.StatusCreated, "success to create article", nil)
}
