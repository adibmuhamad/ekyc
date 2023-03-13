package main

import (
	"id/projects/ekyc/handler"
	"id/projects/ekyc/ocr"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	ocrService := ocr.NewService()

	ocrHandler := handler.NewOcrHandler(ocrService)

	router := gin.Default()
	router.Use(cors.Default())
	
	api := router.Group("/api/v1")

	api.POST("/ktp", ocrHandler.CheckOcrKtp)
	api.POST("/npwp", ocrHandler.CheckOcrNpwp)

	router.Run()

}
