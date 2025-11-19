package service

import (
	"context"
	"hightalent-assessment-task/internal/models"
	"hightalent-assessment-task/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Question interface {
	Create(ctx context.Context, test string) (*models.Question, error)
}

type Service struct {
	Question
}

func NewService(repos repository.Question) *Service {
	return &Service{
		Question: NewQuestionService(repos),
	}
}
