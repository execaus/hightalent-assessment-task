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
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
