package ocr

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"

	// "github.com/anthonynsimon/bild/effect"
	// "github.com/anthonynsimon/bild/segment"
	"github.com/otiai10/gosseract/v2"
)

type service struct {
}

type Service interface {
	CheckOcrKtp(input OcrInput) (OcrKtp, error)
	CheckOcrNpwp(input OcrInput) (OcrNpwp, error)
	CheckOcrSim(input OcrInput) (OcrSim, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) CheckOcrKtp(input OcrInput) (OcrKtp, error) {
	ocr := OcrKtp{}

	// Get base64 from json request
	base64Image := input.OcrImage

	// Decode base64 to byte
	sDec, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return ocr, err
	}

	// Decode byte to image struct
	img, _, err := image.Decode(bytes.NewReader(sDec))
	if err != nil {
		return ocr, err
	}

	// Validate image ktp
	err = ValidateImage(sDec)
	if err != nil {
		return ocr, err
	}

	// Convert Image to grayscale
	// grayscale := effect.Grayscale(img)

	// Convert Image to threshold segment
	// threshold := segment.Threshold(grayscale, 128)

	// Convert Image to Bytes
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

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
		return ocr, err
	}

	// Validate data ktp
	err = ValidateImageKtp(text)
	if err != nil {
		return ocr, err
	}

	ocr, err = FormatDataKtp(text)
	if err != nil {
		return ocr, err
	}

	return ocr, nil

}

func (s *service) CheckOcrSim(input OcrInput) (OcrSim, error) {
	ocr := OcrSim{}

	// Get base64 from json request
	base64Image := input.OcrImage

	// Decode base64 to byte
	sDec, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return ocr, err
	}

	// Decode byte to image struct
	img, _, err := image.Decode(bytes.NewReader(sDec))
	if err != nil {
		return ocr, err
	}

	// Validate image ktp
	err = ValidateImage(sDec)
	if err != nil {
		return ocr, err
	}

	// Convert Image to grayscale
	// grayscale := effect.Grayscale(img)

	// Convert Image to threshold segment
	// threshold := segment.Threshold(grayscale, 128)

	// Convert Image to Bytes
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	// Initiation Gosseract new client
	client := gosseract.NewClient()

	// close client when the main function is finished running
	defer client.Close()

	// Read byte to image and set whitelist character
	client.SetImageFromBytes(buf.Bytes())
	client.SetLanguage("eng")

	// Get text result from OCR
	text, err := client.Text()
	if err != nil {
		return ocr, err
	}

	// Validate data sim
	err = ValidateImageSim(text)
	if err != nil {
		return ocr, err
	}

	ocr, err = FormatDataSim(text)
	if err != nil {
		return ocr, err
	}

	return ocr, nil

}

func (s *service) CheckOcrNpwp(input OcrInput) (OcrNpwp, error) {
	ocr := OcrNpwp{}

	// Get base64 from json request
	base64Image := input.OcrImage

	// Decode base64 to byte
	sDec, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return ocr, err
	}

	// Decode byte to image struct
	img, _, err := image.Decode(bytes.NewReader(sDec))
	if err != nil {
		return ocr, err
	}

	// Validate image ktp
	err = ValidateImage(sDec)
	if err != nil {
		return ocr, err
	}

	// Convert Image to grayscale
	// grayscale := effect.Grayscale(img)

	// Convert Image to threshold segment
	// threshold := segment.Threshold(grayscale, 128)

	// Convert Image to Bytes
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	// Initiation Gosseract new client
	client := gosseract.NewClient()

	// close client when the main function is finished running
	defer client.Close()

	// Read byte to image and set whitelist character
	client.SetImageFromBytes(buf.Bytes())
	client.SetLanguage("eng")

	// Get text result from OCR
	text, err := client.Text()
	if err != nil {
		return ocr, err
	}

	// Validate data npwp
	err = ValidateImageNpwp(text)
	if err != nil {
		return ocr, err
	}

	ocr, err = FormatDataNpwp(text)
	if err != nil {
		return ocr, err
	}

	return ocr, nil

}
