package detect

type FaceInput struct {
	FaceImage string `json:"image" binding:"required"`
}
