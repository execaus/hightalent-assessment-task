package router

import "net/http"

type HTTPError interface {
	error
	StatusCode() int
}

type BadRequestError struct {
	message string
}

func (e BadRequestError) Error() string {
	return e.message
}

func (e BadRequestError) StatusCode() int {
	return http.StatusBadRequest
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{message: message}
}
