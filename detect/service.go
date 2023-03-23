package detect

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math"

	"gocv.io/x/gocv"

	"github.com/Kagami/go-face"
	"github.com/corona10/goimagehash"
	"github.com/nfnt/resize"
	"github.com/otiai10/gosseract/v2"
)

const dataDir = "dataset/face/models"

type service struct {
}

type Service interface {
	DetectFace(input ImageInput) (DetectFace, error)
	CompareFace(input CompareFaceInput) (CompareFace, error)
	CompareSignature(input CompareSignatureInput) (CompareSignature, error)
	ImageForgery(input ImageInput) (ForgeryImage, error)
}

func NewService() *service {
	return &service{}
}

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
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

	dataFace.Valid = true

	return dataFace, nil

}

func (s *service) CompareFace(input CompareFaceInput) (CompareFace, error) {
	// Get base64 from json request
	base64KtpImage := input.KTPImage
	base64FaceImage := input.FaceImage

	// Decode base64 to byte
	sFirstDec, err := base64.StdEncoding.DecodeString(base64KtpImage)
	if err != nil {
		return CompareFace{}, errors.New("Error on ktp image: " + err.Error())
	}

	sSecondDec, err := base64.StdEncoding.DecodeString(base64FaceImage)
	if err != nil {
		return CompareFace{}, errors.New("Error on face image: " + err.Error())
	}

	// Validate image
	err = ValidateImage(sFirstDec)
	if err != nil {
		return CompareFace{}, errors.New("Error on ktp image: " + err.Error())
	}

	err = ValidateImage(sSecondDec)
	if err != nil {
		return CompareFace{}, errors.New("Error on face image: " + err.Error())
	}

	// Decode byte to image struct
	imgOcr, _, err := image.Decode(bytes.NewReader(sFirstDec))
	if err != nil {
		return CompareFace{}, err
	}

	// Convert Image to Bytes
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, imgOcr, nil)

	// Initiation Gosseract new client
	client := gosseract.NewClient()

	// close client when the main function is finished running
	defer client.Close()

	// Read byte to image and set whitelist character
	client.SetImageFromBytes(buf.Bytes())
	client.SetLanguage("eng")
	client.SetVariable("load_system_dawg", "0")
	client.SetVariable("load_freq_dawg", "0")

	// Get text result from OCR
	text, err := client.Text()

	if err != nil {
		return CompareFace{}, err
	}

	// Validate data ktp
	err = ValidateImageKtp(text)
	if err != nil {
		return CompareFace{}, err
	}

	// Init the recognizer
	rec, err := face.NewRecognizer(dataDir)
	if err != nil {
		return CompareFace{}, err
	}

	// Free the resources
	defer rec.Close()

	// Recognize faces on that image.
	ktpFaces, err := rec.RecognizeCNN(sFirstDec)
	if err != nil {
		return CompareFace{}, errors.New("Error on ktp image: " + err.Error())
	}

	otherFaces, err := rec.RecognizeCNN(sSecondDec)
	if err != nil {
		return CompareFace{}, errors.New("Error on face image: " + err.Error())
	}

	dataCompare := CompareFace{}
	dataCompare.Similarity = "0"
	dataCompare.KTPImageValid = len((ktpFaces)) == 0
	dataCompare.FaceImageValid = len((otherFaces)) == 0
	if len((ktpFaces)) == 0 {
		return dataCompare, errors.New("No face detected on ktp image")
	}

	if len((otherFaces)) == 0 {
		return dataCompare, errors.New("No face detected on face image")
	}

	dataCompare.KTPImageValid = len((ktpFaces)) > 1
	dataCompare.FaceImageValid = len((otherFaces)) > 1

	if len((ktpFaces)) > 1 {
		return dataCompare, errors.New("Multiple face detected on ktp image")
	}

	if len((otherFaces)) > 1 {
		return dataCompare, errors.New("Multiple face detected on face image")
	}

	dataCompare.KTPImageValid = true
	dataCompare.FaceImageValid = true

	firstSecondDistance := face.SquaredEuclideanDistance(ktpFaces[0].Descriptor, otherFaces[0].Descriptor)
	distance := fmt.Sprintf("%.8f", (1 - firstSecondDistance))

	dataCompare.Similarity = distance

	return dataCompare, nil
}

func (s *service) CompareSignature(input CompareSignatureInput) (CompareSignature, error) {
	// Get base64 from json request
	base64KtpImage := input.KTPImage
	base64SignatureImage := input.SignatureImage

	// Decode base64 to byte
	sFirstDec, err := base64.StdEncoding.DecodeString(base64KtpImage)
	if err != nil {
		return CompareSignature{}, errors.New("Error on ktp image: " + err.Error())
	}

	sSecondDec, err := base64.StdEncoding.DecodeString(base64SignatureImage)
	if err != nil {
		return CompareSignature{}, errors.New("Error on signature image: " + err.Error())
	}

	// Validate image
	err = ValidateImage(sFirstDec)
	if err != nil {
		return CompareSignature{}, errors.New("Error on ktp image: " + err.Error())
	}

	err = ValidateImage(sSecondDec)
	if err != nil {
		return CompareSignature{}, errors.New("Error on signature image: " + err.Error())
	}

	// Decode byte to image struct
	ktpImg, _, err := image.Decode(bytes.NewReader(sFirstDec))
	if err != nil {
		return CompareSignature{}, err
	}

	signatureImg, _, err := image.Decode(bytes.NewReader(sSecondDec))
	if err != nil {
		return CompareSignature{}, err
	}

	// Convert Image to Bytes
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, ktpImg, nil)

	// Initiation Gosseract new client
	client := gosseract.NewClient()

	// close client when the main function is finished running
	defer client.Close()

	// Read byte to image and set whitelist character
	client.SetImageFromBytes(buf.Bytes())
	client.SetLanguage("eng")
	client.SetVariable("load_system_dawg", "0")
	client.SetVariable("load_freq_dawg", "0")

	// Get text result from OCR
	text, err := client.Text()

	if err != nil {
		return CompareSignature{}, err
	}

	// Validate data ktp
	err = ValidateImageKtp(text)
	if err != nil {
		return CompareSignature{}, err
	}

	// Get the image dimensions
	width := ktpImg.Bounds().Max.X
	height := ktpImg.Bounds().Max.Y

	// top-left corner of the signature area
	x0 := int(float64(width) * 0.7)
	y0 := int(float64(height) * 0.8)
	x1, y1 := width, height // bottom-right corner of the signature area

	tempImg := ktpImg.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(x0, y0, x1, y1))

	// Convert Image to Bytes
	buf = new(bytes.Buffer)
	jpeg.Encode(buf, tempImg, nil)

	// Convert the image to grayscale
	grayKtp := image.NewGray(tempImg.Bounds())

	for x := grayKtp.Rect.Min.X; x < grayKtp.Rect.Max.X; x++ {
		for y := grayKtp.Rect.Min.Y; y < grayKtp.Rect.Max.Y; y++ {
			// Get the color of the current pixel
			c := color.GrayModel.Convert(tempImg.At(x, y)).(color.Gray)

			// Set the gray color value for the current pixel
			grayKtp.Set(x, y, c)
		}
	}

	graySignature := image.NewGray(signatureImg.Bounds())

	for x := graySignature.Rect.Min.X; x < graySignature.Rect.Max.X; x++ {
		for y := graySignature.Rect.Min.Y; y < graySignature.Rect.Max.Y; y++ {
			// Get the color of the current pixel
			c := color.GrayModel.Convert(signatureImg.At(x, y)).(color.Gray)

			// Set the gray color value for the current pixel
			graySignature.Set(x, y, c)
		}
	}

	// Resize the image
	width = 200
	height = 200
	resizedKtp := resize.Resize(uint(width), uint(height), grayKtp, resize.Lanczos3)
	resizedSignature := resize.Resize(uint(width), uint(height), graySignature, resize.Lanczos3)

	// Compute the perceptual hash of the images
	hash1, err := goimagehash.PerceptionHash(resizedKtp)
	if err != nil {
		return CompareSignature{}, errors.New("Error on ktp image: " + err.Error())
	}

	hash2, err := goimagehash.PerceptionHash(resizedSignature)
	if err != nil {
		return CompareSignature{}, errors.New("Error on signature image: " + err.Error())
	}

	// Compute the hamming distance between the hashes
	distance, err := hash1.Distance(hash2)
	if err != nil {
		return CompareSignature{}, err
	}

	// Print the similarity score
	similarity := 1.0 - (float64(distance) / float64(hash1.Bits()))

	dataCompare := CompareSignature{}
	dataCompare.KTPImageValid = true
	dataCompare.SignatureImageValid = true

	dataCompare.Similarity = fmt.Sprintf("%.8f", similarity)

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

	img, err := gocv.IMDecode(sDec, gocv.IMReadColor)
	if err != nil {
		return ForgeryImage{}, errors.New("Error on signature image: " + err.Error())
	}

	// calculate the average pixel value in each block of the image
	blockWidth := 16
	blockHeight := 16
	rows := int(math.Floor(float64(img.Rows() / blockHeight)))
	cols := int(math.Floor(float64(img.Cols() / blockWidth)))
	var blockAvgVals []float64
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			block := img.Region(image.Rect(j*blockWidth, i*blockHeight, blockWidth, blockHeight))
			blockAvg := block.Mean()
			blockAvgVals = append(blockAvgVals, blockAvg.Val1)
		}
	}

	// calculate the standard deviation of the block pixel values
	sum := 0.0
	for _, v := range blockAvgVals {
		sum += v
	}
	mean := sum / float64(len(blockAvgVals))

	var variance float64
	for _, v := range blockAvgVals {
		variance += (v - mean) * (v - mean)
	}
	variance = variance / float64(len(blockAvgVals)-1)

	precision := math.Sqrt(variance)

	formater := ForgeryImage{}
	if precision > 50.0 {
		formater.Forged = true
	} else {
		formater.Forged = false
	}

	formater.Precision = fmt.Sprintf("%.8f", (precision / 100))

	return formater, nil

}
