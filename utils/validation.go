package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateRequest(c *gin.Context, input interface{}) error {
	if err := c.ShouldBindJSON(input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{Field: fe.Field(), Message: getErrorMsg(fe)}
			}
			JSONError(c, http.StatusBadRequest, "Validation failed", out)
		} else {
			JSONError(c, http.StatusBadRequest, "Invalid request body", nil)
		}
		return err
	}
	return nil
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	default:
		return "Invalid input"
	}
}
