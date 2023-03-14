package verify

import (
	"errors"

	email "github.com/cameronnewman/go-emailvalidation/v3"
)

func ValidateEmail(text string) (err error) {
	if len(text) < 5 {
		return errors.New("Invalid Email")
	}

	err = email.Validate(text)
	if err != nil {
		return err
	}

	return nil
}
