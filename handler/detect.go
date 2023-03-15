package handler

import (
	"id/projects/ekyc/detect"
	"id/projects/ekyc/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type detectHandler struct {
	detectService detect.Service
}

func NewDetectHandler(detectService detect.Service) *detectHandler {
	return &detectHandler{detectService}
}

func (h *detectHandler) DetectFace(c *gin.Context) {
	var input detect.FaceInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Unable to process request", http.StatusUnprocessableEntity, "FAILED", errorMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	newDetect, err := h.detectService.DetectFace(input)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "FAILED", newDetect)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("OK", http.StatusOK, "SUCCESS", newDetect)
	c.JSON(http.StatusOK, response)

}

