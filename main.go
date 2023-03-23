package main

import (
	"id/projects/ekyc/detect"
	"id/projects/ekyc/handler"
	"id/projects/ekyc/ocr"
	"id/projects/ekyc/parser"
	"id/projects/ekyc/verify"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	ocrService := ocr.NewService()
	detectService := detect.NewService()
	parserService := parser.NewService()
	verifyService := verify.NewService()

	ocrHandler := handler.NewOcrHandler(ocrService)
	detectHandler := handler.NewDetectHandler(detectService)
	parserHandler := handler.NewParserHandler(parserService)
	verifyHandler := handler.NewVerifyHandler(verifyService)

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")

	api.POST("/ocr/ktp", ocrHandler.CheckOcrKtp)
	api.POST("/ocr/npwp", ocrHandler.CheckOcrNpwp)
	api.POST("/ocr/sim", ocrHandler.CheckOcrSim)
	api.POST("/ocr/passport", ocrHandler.CheckOcrPassport)

	api.POST("/detect/face-detection", detectHandler.DetectFace)
	api.POST("/detect/face-compare", detectHandler.CompareFace)
	api.POST("/detect/signature-compare", detectHandler.CompareSignature)
	api.POST("/detect/forgery", detectHandler.ImageForgery)

	api.POST("/parser/ktp", parserHandler.ParserDataNik)
	api.POST("/parser/npwp", parserHandler.ParserDataNpwp)
	api.POST("/parser/sim", parserHandler.ParserDataSim)

	api.POST("/verify/email", verifyHandler.VerifyEmail)

	router.Run()

}
