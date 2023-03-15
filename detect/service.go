package detect

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"

	"id/projects/ekyc/lib/forgery"

	"github.com/Kagami/go-face"
)

const dataDir = "dataset/face/models"

type service struct {
}

type Service interface {
	DetectFace(input ImageInput) (DetectFace, error)
	CompareFace(input CompareInput) (CompareFace, error)
	ImageForgery(input ImageInput) (ForgeryImage, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) DetectFace(input ImageInput) (DetectFace, error) {
	// Get base64 from json request
	base64Image := input.Image

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
	faces, err := rec.RecognizeCNN(sDec)
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

	// Decode byte to image struct
	// img, _, err := image.Decode(bytes.NewReader(sDec))
	// if err != nil {
	// 	return dataFace, err
	// }

	// precision := float64(forgery(img))
	// if precision > 50.0 {
	// 	print("%.8f the image is forged!", precision)
	// } else {
	// 	precision = 100 - precision
	// 	print("%.8f the image is NOT forged!", precision)
	// }

	dataFace.Valid = true

	return dataFace, nil

}

func (s *service) CompareFace(input CompareInput) (CompareFace, error) {
	// Get base64 from json request
	base64ImageFirst := input.FirstImage
	base64ImageSecond := input.SecondImage

	// Decode base64 to byte
	sFirstDec, err := base64.StdEncoding.DecodeString(base64ImageFirst)
	if err != nil {
		return CompareFace{}, errors.New("Error on first image: " + err.Error())
	}

	sSecondDec, err := base64.StdEncoding.DecodeString(base64ImageSecond)
	if err != nil {
		return CompareFace{}, errors.New("Error on second image: " + err.Error())
	}

	// Validate image
	err = ValidateImage(sFirstDec)
	if err != nil {
		return CompareFace{}, errors.New("Error on first image: " + err.Error())
	}

	err = ValidateImage(sSecondDec)
	if err != nil {
		return CompareFace{}, errors.New("Error on second image: " + err.Error())
	}

	// Init the recognizer
	rec, err := face.NewRecognizer(dataDir)
	if err != nil {
		return CompareFace{}, err
	}

	// Free the resources
	defer rec.Close()

	// Recognize faces on that image.
	firstFaces, err := rec.RecognizeCNN(sFirstDec)
	if err != nil {
		return CompareFace{}, errors.New("Error on first image: " + err.Error())
	}

	secondFaces, err := rec.RecognizeCNN(sSecondDec)
	if err != nil {
		return CompareFace{}, errors.New("Error on second image: " + err.Error())
	}

	dataCompare := CompareFace{}
	dataCompare.Similarity = "0"
	dataCompare.FirstImageValid = len((firstFaces)) == 0
	dataCompare.SecondImageValid = len((secondFaces)) == 0
	if len((firstFaces)) == 0 {
		return dataCompare, errors.New("No face detected on first image")
	}

	if len((secondFaces)) == 0 {
		return dataCompare, errors.New("No face detected on second image")
	}

	dataCompare.FirstImageValid = len((firstFaces)) > 1
	dataCompare.SecondImageValid = len((secondFaces)) > 1

	if len((firstFaces)) > 1 {
		return dataCompare, errors.New("Multiple face detected on first image")
	}

	if len((secondFaces)) > 1 {
		return dataCompare, errors.New("Multiple face detected on second image")
	}

	dataCompare.FirstImageValid = true
	dataCompare.SecondImageValid = true

	firstSecondDistance := face.SquaredEuclideanDistance(firstFaces[0].Descriptor, secondFaces[0].Descriptor)
	distance := fmt.Sprintf("%.8f", (1 - firstSecondDistance))

	dataCompare.Similarity = distance

	return dataCompare, nil

}

func (s *service) ImageForgery(input ImageInput) (ForgeryImage, error) {
	// Get base64 from json request
	base64Image := input.Image

	// Decode base64 to byte
	sDec, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return ForgeryImage{}, err
	}

	// Validate image
	err = ValidateImage(sDec)
	if err != nil {
		return ForgeryImage{}, err
	}

	// Decode byte to image struct
	img, _, err := image.Decode(bytes.NewReader(sDec))
	if err != nil {
		return ForgeryImage{}, err
	}

	precision := float64(forgery.Forgery(img))

	formater := ForgeryImage{}
	if precision > 50.0 {
		formater.Forged = true
		formater.Precision = fmt.Sprintf("%.8f", precision)
	} else {
		formater.Forged = false
		formater.Precision = fmt.Sprintf("%.8f", precision)
	}

	return formater, nil

}
