package repository

import "hightalent-assessment-task/internal/models"

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

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
