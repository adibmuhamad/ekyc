package handler

import (
	"id/projects/ekyc/helper"
	"id/projects/ekyc/verify"
	"net/http"

	"github.com/gin-gonic/gin"
)

type verifyHandler struct {
	verifyService verify.Service
}

func NewVerifyHandler(verifyService verify.Service) *verifyHandler {
	return &verifyHandler{verifyService}
}

func (h *verifyHandler) VerifyEmail(c *gin.Context) {
	var input verify.EmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Unable to process request", http.StatusUnprocessableEntity, "FAILED", errorMessage)
		c.JSON(http.StatusOK, response)
		return
	}

	newVerify, err := h.verifyService.VerifyEmail(input)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "FAILED", newVerify)
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.APIResponse("OK", http.StatusOK, "SUCCESS", newVerify)
	c.JSON(http.StatusOK, response)

}
