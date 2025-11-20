package models

import (
	"time"
)

type Question struct {
	ID        uint      `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateQuestionRequest struct {
	Text string `json:"text" validate:"required,min=1"`
}

type CreateQuestionResponse struct {
	Question *Question `json:"question"`
}

type GetAllQuestionsResponse struct {
	Questions []*Question `json:"questions"`
}

type GetQuestionResponse struct {
	Question *Question `json:"question"`
	Answers  []*Answer `json:"answers"`
}

type DeleteQuestionResponse struct {
	Question *Question `json:"question"`
}
