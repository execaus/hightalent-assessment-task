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
	Text       string `json:"text"`
	QuestionID uint   `json:"question_id"`
}

type CreateAnswerResponse struct {
	Answer *Answer `json:"answer"`
}

type GetAnswerResponse struct {
	Answer *Answer `json:"answer"`
}
