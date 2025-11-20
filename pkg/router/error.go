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

type UnauthorizedError struct {
	message string
}

func (e UnauthorizedError) Error() string {
	return e.message
}

func (e UnauthorizedError) StatusCode() int {
	return http.StatusUnauthorized
}

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{
		message: message,
	}
}

type BusinessLoginError struct {
	message string
}

func (e BusinessLoginError) Error() string {
	return e.message
}

func NewBusinessLogicError(message string) *BusinessLoginError {
	return &BusinessLoginError{message: message}
}
