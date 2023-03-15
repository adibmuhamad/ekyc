package ocr

import (
	"bytes"
	"errors"
	"image"
	"regexp"
	"strings"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func ValidateImage(data []byte, format string) (err error) {
	images, _, err := image.DecodeConfig(bytes.NewReader(data))

	if err != nil {
		return err
	}

	if images.Height >= images.Width {
		return errors.New("Images must be landscape")
	}

	if images.Height < 256 {
		return errors.New("Minimum resolution of 256 x 256")
	}

	if images.Height > 4096 {
		return errors.New("Maximum resolution of 4096 x 4096")
	}

	if format == "jpeg" || format == "png" {
		return nil
	}

	return errors.New("Image format is not jpeg/png")

}

func ValidateImageKtp(data string) (err error) {
	if len(strings.TrimSpace(data)) == 0 || err != nil {
		return errors.New("Invalid KTP image")
	}

	var re = regexp.MustCompile(`NIK`)
	var nonRe = regexp.MustCompile(`NPWP`)

	if re.MatchString(data) && !nonRe.MatchString(data) {
		return nil
	}

	return errors.New("Invalid KTP image")
}

func ValidateImageNpwp(data string) (err error) {
	if len(strings.TrimSpace(data)) == 0 || err != nil {
		return errors.New("Invalid KTP image")
	}

	var re = regexp.MustCompile(`NPWP`)

	if re.MatchString(data) {
		return nil
	}

	return errors.New("Invalid NPWP image")
}

func ValidateImageSim(data string) (err error) {
	if len(strings.TrimSpace(data)) == 0 || err != nil {
		return errors.New("Invalid SIM image")
	}

	var re = regexp.MustCompile(`SURAT IZIN MENGEMUDI`)

	if re.MatchString(data) {
		return nil
	}

	return errors.New("Invalid SIM image")
}

func ValidateImagePassport(data string) (err error) {
	if len(strings.TrimSpace(data)) == 0 || err != nil {
		return errors.New("Invalid Passport image")
	}

	var re = regexp.MustCompile(`REPUBLIC OF INDONESIA`)

	if re.MatchString(data) {
		return nil
	}

	return errors.New("Invalid Passport image")
}
