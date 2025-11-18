package service

import (
	"hightalent-assessment-task/internal/models"
	"hightalent-assessment-task/internal/repository"
)

type Question interface {
	Create(test string) (*models.Question, error)
}

type Service struct {
	Question
}

func NewService(repos repository.Repository) *Service {
	return &Service{
		Question: NewQuestionService(repos),
	}
}
