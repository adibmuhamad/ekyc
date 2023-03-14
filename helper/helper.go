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
		return "JANUARY"
	case 2:
		return "FEBRUARY"
	case 3:
		return "MARCH"
	case 4:
		return "APRIL"
	case 5:
		return "MAY"
	case 6:
		return "JUNE"
	case 7:
		return "JULY"
	case 8:
		return "AUGUST"
	case 9:
		return "SEPTEMBER"
	case 10:
		return "OCTOBER"
	case 11:
		return "NOVEMBER"
	case 12:
		return "DECEMBER"
	default:
		return ""
	}
}
