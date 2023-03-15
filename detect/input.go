package detect

type FaceInput struct {
	FaceImage string `json:"image" binding:"required"`
}

type CompareInput struct {
	FirstImage  string `json:"firstImage" binding:"required"`
	SecondImage string `json:"secondImage" binding:"required"`
}
