package handler

import (
	"id/projects/ekyc/helper"
	"id/projects/ekyc/ocr"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ocrHandler struct {
	ocrService ocr.Service
}

func NewOcrHandler(ocrService ocr.Service) *ocrHandler {
	return &ocrHandler{ocrService}
}

func (h *ocrHandler) CheckOcrKtp(c *gin.Context) {
	var input ocr.OcrInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Unable to process request", http.StatusUnprocessableEntity, "FAILED", errorMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	newOcr, err := h.ocrService.CheckOcrKtp(input)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "FAILED", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("OK", http.StatusOK, "SUCCESS", newOcr)
	c.JSON(http.StatusOK, response)

}

func (h *ocrHandler) CheckOcrNpwp(c *gin.Context) {
	var input ocr.OcrInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Unable to process request", http.StatusUnprocessableEntity, "FAILED", errorMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	newOcr, err := h.ocrService.CheckOcrNpwp(input)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "FAILED", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("OK", http.StatusOK, "SUCCESS", newOcr)
	c.JSON(http.StatusOK, response)

}

func (h *ocrHandler) CheckOcrSim(c *gin.Context) {
	var input ocr.OcrInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Unable to process request", http.StatusUnprocessableEntity, "FAILED", errorMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	newOcr, err := h.ocrService.CheckOcrSim(input)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "FAILED", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("OK", http.StatusOK, "SUCCESS", newOcr)
	c.JSON(http.StatusOK, response)

}

func (h *ocrHandler) CheckOcrPassport(c *gin.Context) {
	var input ocr.OcrInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Unable to process request", http.StatusUnprocessableEntity, "FAILED", errorMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	newOcr, err := h.ocrService.CheckOcrPassport(input)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "FAILED", nil)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("OK", http.StatusOK, "SUCCESS", newOcr)
	c.JSON(http.StatusOK, response)

}
