package repository

import (
	"hightalent-assessment-task/internal/models"
	"time"

	"github.com/google/uuid"
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

type UserTable struct {
	ID       uuid.UUID `gorm:"type:uuid"`
	Login    string
	Password string
}

func (t *UserTable) TableName() string {
	return "app.users"
}

func (t *UserTable) ToModel() *models.User {
	return &models.User{
		ID:       t.ID,
		Login:    t.Login,
		Password: t.Password,
	}
}

type AnswerTable struct {
	ID         uint
	QuestionID uint
	UserID     uuid.UUID `gorm:"type:uuid"`
	Text       string
	CreatedAt  time.Time
}

func (t *AnswerTable) TableName() string {
	return "app.answers"
}

func (t *AnswerTable) ToModel() *models.Answer {
	return &models.Answer{
		ID:         t.ID,
		Text:       t.Text,
		QuestionID: t.QuestionID,
		UserID:     t.UserID,
		CreatedAt:  t.CreatedAt,
	}
}
