package models

import (
	"time"

	"github.com/google/uuid"
)

type Answer struct {
	ID         uint      `json:"id"`
	Text       string    `json:"text"`
	QuestionID uint      `json:"question_id"`
	UserID     uuid.UUID `json:"-"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateAnswerRequest struct {
	Text       string `json:"text" validate:"required,min=1"`
	QuestionID uint   `json:"question_id" validate:"required"`
}

type CreateAnswerResponse struct {
	Answer *Answer `json:"answer"`
}

type GetAnswerResponse struct {
	Answer *Answer `json:"answer"`
}

type DeleteAnswerResponse struct {
	Answer *Answer `json:"answer"`
}
