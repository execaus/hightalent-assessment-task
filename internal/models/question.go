package models

import "time"

type Question struct {
	ID        int
	Text      string
	CreatedAt time.Time
}

type CreateQuestionRequest struct {
	Text string `json:"text"`
}

type CreateQuestionResponse struct {
	Question *Question `json:"question"`
}
