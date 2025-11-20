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
	Text string `json:"text"`
}

type CreateQuestionResponse struct {
	Question *Question `json:"question"`
}

type GetAllQuestionsResponse struct {
	Questions []*Question `json:"questions"`
}
