package repository

import "time"

type QuestionTable struct {
	ID        uint
	Text      string
	CreatedAt time.Time
}

func (t *QuestionTable) TableName() string {
	return "app.questions"
}
