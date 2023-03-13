package ocr

type OcrInput struct {
	OcrImage string `json:"image" binding:"required"`
}
