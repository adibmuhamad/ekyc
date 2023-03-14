package handler

import (
	"id/projects/ekyc/helper"
	"id/projects/ekyc/parser"
	"net/http"

	"github.com/gin-gonic/gin"
)

type parserHandler struct {
	parserService parser.Service
}

func NewParserHandler(parserService parser.Service) *parserHandler {
	return &parserHandler{parserService}
}

func (h *parserHandler) ParserDataNik(c *gin.Context) {
	var input parser.ParserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Unable to process request", http.StatusUnprocessableEntity, "FAILED", errorMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	newOcr, err := h.parserService.ParseDataKtp(input)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "FAILED", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("OK", http.StatusOK, "SUCCESS", newOcr)
	c.JSON(http.StatusOK, response)

}
