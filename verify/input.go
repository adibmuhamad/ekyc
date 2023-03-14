package verify

type EmailInput struct {
	Email string `json:"email" binding:"required"`
}
