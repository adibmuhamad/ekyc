package detect

import (
	"encoding/base64"
	"errors"

	"github.com/Kagami/go-face"
)

const dataDir = "dataset/face/models"

type service struct {
}

type Service interface {
	DetectFace(input FaceInput) (DetectFace, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) DetectFace(input FaceInput) (DetectFace, error) {

	// Get base64 from json request
	base64Image := input.FaceImage

	// Decode base64 to byte
	sDec, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return DetectFace{}, err
	}

	// Validate image
	err = ValidateImage(sDec)
	if err != nil {
		return DetectFace{}, err
	}

	// Init the recognizer
	rec, err := face.NewRecognizer(dataDir)
	if err != nil {
		return DetectFace{}, err
	}

	// Free the resources
	defer rec.Close()

	// Recognize faces on that image.
	faces, err := rec.Recognize(sDec)
	if err != nil {
		return DetectFace{}, err
	}

	dataFace := DetectFace{}
	if len((faces)) == 0 {
		dataFace.Valid = false
		return dataFace, errors.New("No face detected")
	}

	if len((faces)) > 1 {
		dataFace.Valid = false
		return dataFace, errors.New("Multiple face detected")
	}

	dataFace.Valid = true

	return dataFace, nil

}
