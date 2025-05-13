package errors

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
)

type ErrorType string

const (
	ErrorTypeValidation   ErrorType = "VALIDATION_ERROR"
	ErrorTypeNotFound     ErrorType = "NOT_FOUND"
	ErrorTypeUnauthorized ErrorType = "UNAUTHORIZED"
	ErrorTypeForbidden    ErrorType = "FORBIDDEN"
	ErrorTypeInternal     ErrorType = "INTERNAL_ERROR"
	ErrorTypeBadRequest   ErrorType = "BAD_REQUEST"
	ErrorTypeConflict     ErrorType = "CONFLICT"
)

type Error struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
	Err     error     `json:"-"`
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %v", e.Type, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) StatusCode() int {
	switch e.Type {
	case ErrorTypeValidation:
		return http.StatusBadRequest
	case ErrorTypeNotFound:
		return http.StatusNotFound
	case ErrorTypeUnauthorized:
		return http.StatusUnauthorized
	case ErrorTypeForbidden:
		return http.StatusForbidden
	case ErrorTypeInternal:
		return http.StatusInternalServerError
	case ErrorTypeBadRequest:
		return http.StatusBadRequest
	case ErrorTypeConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func New(errType ErrorType, message string, err error) *Error {
	return &Error{
		Type:    errType,
		Message: message,
		Err:     err,
	}
}

func ValidationError(message string, err error) *Error {
	return New(ErrorTypeValidation, message, err)
}

func NotFoundError(entity string, err error) *Error {
	return New(ErrorTypeNotFound, fmt.Sprintf("%s not found", entity), err)
}

func UnauthorizedError(message string, err error) *Error {
	return New(ErrorTypeUnauthorized, message, err)
}

func ForbiddenError(message string, err error) *Error {
	return New(ErrorTypeForbidden, message, err)
}

func InternalError(message string, err error) *Error {
	return New(ErrorTypeInternal, message, err)
}

func BadRequestError(message string, err error) *Error {
	return New(ErrorTypeBadRequest, message, err)
}

func ConflictError(message string, err error) *Error {
	return New(ErrorTypeConflict, message, err)
}

func AsError(err error) *Error {
	if err == nil {
		return nil
	}

	if e, ok := err.(*Error); ok {
		return e
	}

	return InternalError("unexpected error", err)
}

func RecoverPanic(source string) {
	if r := recover(); r != nil {
		stackTrace := debug.Stack()

		slog.Error("Exception caught", "error", r, "source", source, "stack_trace", string(stackTrace))
	}
}

func RecoverPanicWithCallback(source string, callback func(err interface{}, stack []byte)) {
	if r := recover(); r != nil {
		stackTrace := debug.Stack()

		slog.Error("Exception caught", "error", r, "source", source, "stack_trace", string(stackTrace))

		if callback != nil {
			callback(r, stackTrace)
		}
	}
}
