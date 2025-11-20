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
