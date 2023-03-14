package main

import (
	"id/projects/ekyc/handler"
	"id/projects/ekyc/ocr"
	"id/projects/ekyc/parser"
	"id/projects/ekyc/verify"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	ocrService := ocr.NewService()
	parserService := parser.NewService()
	verifyService := verify.NewService()

	ocrHandler := handler.NewOcrHandler(ocrService)
	parserHandler := handler.NewParserHandler(parserService)
	verifyHandler := handler.NewVerifyHandler(verifyService)

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")

	api.POST("/ocr/ktp", ocrHandler.CheckOcrKtp)
	api.POST("/ocr/npwp", ocrHandler.CheckOcrNpwp)
	api.POST("/ocr/sim", ocrHandler.CheckOcrSim)
	api.POST("/ocr/passport", ocrHandler.CheckOcrPassport)

	api.POST("/parser/ktp", parserHandler.ParserDataNik)
	api.POST("/parser/npwp", parserHandler.ParserDataNpwp)
	api.POST("/parser/sim", parserHandler.ParserDataSim)

	api.POST("/verify/email", verifyHandler.VerifyEmail)

	router.Run()

}
