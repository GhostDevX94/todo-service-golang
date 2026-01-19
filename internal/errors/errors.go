package errors

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound     = errors.New("resource not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
	ErrValidation   = errors.New("validation error")
	ErrInternal     = errors.New("internal server error")
	ErrConflict     = errors.New("resource conflict")
	ErrBadRequest   = errors.New("bad request")
)

type AppError struct {
	Err     error
	Message string
	Code    int
}

func (e *AppError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewAppError(err error, message string, code int) *AppError {
	return &AppError{
		Err:     err,
		Message: message,
		Code:    code,
	}
}

func NotFound(message string) *AppError {
	return &AppError{
		Err:     ErrNotFound,
		Message: message,
		Code:    404,
	}
}

func Unauthorized(message string) *AppError {
	return &AppError{
		Err:     ErrUnauthorized,
		Message: message,
		Code:    401,
	}
}

func Forbidden(message string) *AppError {
	return &AppError{
		Err:     ErrForbidden,
		Message: message,
		Code:    403,
	}
}

func Validation(message string) *AppError {
	return &AppError{
		Err:     ErrValidation,
		Message: message,
		Code:    400,
	}
}

func Internal(message string) *AppError {
	return &AppError{
		Err:     ErrInternal,
		Message: message,
		Code:    500,
	}
}

func Conflict(message string) *AppError {
	return &AppError{
		Err:     ErrConflict,
		Message: message,
		Code:    409,
	}
}

func BadRequest(message string) *AppError {
	return &AppError{
		Err:     ErrBadRequest,
		Message: message,
		Code:    400,
	}
}
