package main

import (
	"id/projects/ekyc/handler"
	"id/projects/ekyc/ocr"
	"id/projects/ekyc/parser"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	ocrService := ocr.NewService()
	parserService := parser.NewService()

	ocrHandler := handler.NewOcrHandler(ocrService)
	parserHandler := handler.NewParserHandler(parserService)

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")

	api.POST("/ocr/ktp", ocrHandler.CheckOcrKtp)
	api.POST("/ocr/npwp", ocrHandler.CheckOcrNpwp)

	api.POST("/parser/ktp", parserHandler.ParserDataNik)

	router.Run()

}
