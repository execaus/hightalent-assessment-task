package service

type BusinessLoginError struct {
	message string
}

func (e BusinessLoginError) Error() string {
	return e.message
}

func NewBusinessLoginError(message string) *BusinessLoginError {
	return &BusinessLoginError{message: message}
}
