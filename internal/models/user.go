package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Login    string    `json:"login"`
	Password string    `json:"-"`
}

type SignUpRequest struct {
	Login    string `json:"login" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=6"`
}

type SignUpResponse struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
