package detect

import (
)

type DetectFace struct {
	Valid bool `json:"valid"`
}

type CompareFace struct {
	KTPImageValid bool `json:"ktpImageValid"`
	FaceImageValid bool `json:"faceImageValid"`
	Similarity string `json:"similarity"`
}

type CompareSignature struct {
	KTPImageValid bool `json:"ktpImageValid"`
	SignatureImageValid bool `json:"signatureImageValid"`
	Similarity string `json:"similarity"`
}

type ForgeryImage struct {
	Forged bool `json:"forged"`
	Precision string `json:"precision"`
}
