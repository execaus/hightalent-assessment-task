package repository

import (
	"hightalent-assessment-task/internal/models"
	"time"
)

type QuestionTable struct {
	ID        uint
	Text      string
	CreatedAt time.Time
}

func (t *QuestionTable) TableName() string {
	return "app.questions"
}

func (t *QuestionTable) ToModel() *models.Question {
	return &models.Question{
		ID:        t.ID,
		Text:      t.Text,
		CreatedAt: t.CreatedAt,
	}
}
