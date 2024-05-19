package api

import "net/http"

func NewUnauthorizedApiError(err error) *Error {
	return &Error{
		Message: "Unauthorized",
		Code:    http.StatusUnauthorized,
		Cause:   err,
	}
}

func NewBadRequestError(err error) *Error {
	return &Error{
		Message: "Request body invalid",
		Code:    http.StatusBadRequest,
		Cause:   err,
	}
}
