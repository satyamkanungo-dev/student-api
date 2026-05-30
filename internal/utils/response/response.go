package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusError = "Error"
)

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) ErrorResponse {
	return ErrorResponse{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) ErrorResponse {
	var errMgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMgs = append(errMgs, fmt.Sprintf("field %s is required field", err.Field()))
		default:
			errMgs = append(errMgs, fmt.Sprintf("field %s is invalid", err.Field()))
		}
	}

	return ErrorResponse{
		Status: StatusError,
		Error:  strings.Join(errMgs, ", "),
	}

}
