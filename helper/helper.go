package helper

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"

)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func Marshaler(a interface{}) (map[string]interface{}, error) {
	marshal, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	var b map[string]interface{}
	err = json.Unmarshal(marshal, &b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func ConvertMonth(b int) string {
	switch b {
	case 1:
		return "January"
	case 2:
		return "February"
	case 3:
		return "March"
	case 4:
		return "April"
	case 5:
		return "May"
	case 6:
		return "June"
	case 7:
		return "July"
	case 8:
		return "August"
	case 9:
		return "September"
	case 10:
		return "October"
	case 11:
		return "November"
	case 12:
		return "December"
	default:
		return ""
	}
}
