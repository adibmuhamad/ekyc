package verify

type service struct {
}

type Service interface {
	VerifyEmail(input EmailInput) (VerifyEmail, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) VerifyEmail(input EmailInput) (VerifyEmail, error) {
	email := VerifyEmail{}
	email.Email = input.Email

	// Validate Email
	err := ValidateEmail(input.Email)
	if err != nil {
		email.Valid = false
		return email, err
	}

	email.Valid = true
	return email, nil

}
