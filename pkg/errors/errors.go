package errors

import "fmt"

// AppError represents a custom application error.
type AppError struct {
	Code    int
	Message string
}

// New creates a new AppError.
func New(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// Error implements the error interface.
func (e *AppError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
