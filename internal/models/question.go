package models

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	ID        uint
	Text      string
	CreatedAt time.Time
}

type QuestionTable struct {
	gorm.Model
	Text string `gorm:"column:text"`
}

type CreateQuestionRequest struct {
	Text string `json:"text"`
}

type CreateQuestionResponse struct {
	Question *Question `json:"question"`
}
