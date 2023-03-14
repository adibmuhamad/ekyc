package parser

type ParserInput struct {
	NumberID string `json:"numberId" binding:"required"`
}
