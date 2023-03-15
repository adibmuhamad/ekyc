package detect

import (
)

type DetectFace struct {
	Valid bool `json:"valid"`
}

type CompareFace struct {
	FirstImageValid bool `json:"firstImageValid"`
	SecondImageValid bool `json:"secondImageValid"`
	Similarity string `json:"similarity"`
}

type ForgeryImage struct {
	Forged bool `json:"forged"`
	Precision string `json:"precision"`
}
