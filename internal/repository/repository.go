package repository

import "hightalent-assessment-task/internal/models"

type Question interface {
	Create(text string) (*models.Question, error)
}

type Repository struct {
	Question
}

func NewRepository() *Repository {
	return &Repository{
		Question: NewQuestionRepository(),
	}
}
