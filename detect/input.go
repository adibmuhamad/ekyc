package detect

type ImageInput struct {
	Image string `json:"image" binding:"required"`
}

type CompareInput struct {
	FirstImage  string `json:"firstImage" binding:"required"`
	SecondImage string `json:"secondImage" binding:"required"`
}

type CompareFaceInput struct {
	KTPImage  string `json:"ktpImage" binding:"required"`
	FaceImage string `json:"faceImage" binding:"required"`
}

type CompareSignatureInput struct {
	KTPImage       string `json:"ktpImage" binding:"required"`
	SignatureImage string `json:"signatureImage" binding:"required"`
}
