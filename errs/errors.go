package errs

import "net/http"

type AppError struct {
	Code int `json:"code,omitempty"`
	Message string `json:"message"`
}

// AsMessage returns only the message of an error
func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

// NewNotFoundError returns a 404 error
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

// NewUnexpectedError returns a 500 server error
func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

// NewValidationError returns a 422 validation error
func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}