package terrors

import (
	"fmt"
	"net/http"
)

type Error interface {
	IsTError()
	Error() string
	GetCode() int
	GetPublicMessage() string
	GetPrivateMessage() string
	GetData() any
}

type BaseErrorSt struct {
	Code           int
	PublicMessage  string
	PrivateMessage string
	Data           any
}

func (e BaseErrorSt) IsTError() {}

func (e BaseErrorSt) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.PublicMessage)
}

func (e BaseErrorSt) GetCode() int {
	return e.Code
}

func (e BaseErrorSt) GetPublicMessage() string {
	return e.PublicMessage
}

func (e BaseErrorSt) GetPrivateMessage() string {
	return e.PrivateMessage
}

func (e BaseErrorSt) GetData() any {
	return e.Data
}

type PrivateError struct {
	BaseErrorSt
}

func NewPrivateError(privateMessage string) *PrivateError {
	return &PrivateError{
		BaseErrorSt{
			http.StatusInternalServerError,
			"Internal error",
			privateMessage,
			nil,
		},
	}
}

type PublicError struct {
	BaseErrorSt
}

func NewPublicError(code int, publicMessage string, privateMessage string, data any) PublicError {
	return PublicError{
		BaseErrorSt{
			code,
			publicMessage,
			privateMessage,
			data,
		},
	}
}

func NewValidationError(publicMessage string, data any) PublicError {
	return PublicError{
		BaseErrorSt{
			http.StatusBadRequest,
			publicMessage,
			publicMessage,
			data,
		},
	}
}

func NewForbiddenError(publicMessage string, data any) PublicError {
	return PublicError{
		BaseErrorSt{
			http.StatusForbidden,
			publicMessage,
			publicMessage,
			data,
		},
	}
}

func NewUnauthorizedError(publicMessage string, data any) PublicError {
	return PublicError{
		BaseErrorSt{
			http.StatusUnauthorized,
			publicMessage,
			publicMessage,
			data,
		},
	}
}

func NewNotFoundError(publicMessage string, data any) PublicError {
	return PublicError{
		BaseErrorSt{
			http.StatusNotFound,
			publicMessage,
			publicMessage,
			data,
		},
	}
}

func IgnoreError[V any](value V, err error) V {
	return value
}

func NewTimeoutError(publicMessage string, data any) PublicError {
	return PublicError{
		BaseErrorSt{
			http.StatusNotFound,
			publicMessage,
			publicMessage,
			data,
		},
	}
}
