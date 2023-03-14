package verify

import (
)

type VerifyEmail struct {
	Email  string `json:"email"`
	Valid bool   `json:"valid"`
}