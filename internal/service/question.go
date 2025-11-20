package service

import (
	"context"
	"hightalent-assessment-task/internal/models"
	"hightalent-assessment-task/internal/repository"
)

type QuestionService struct {
	service    *Service
	repository repository.Question
}

func (s *QuestionService) Create(ctx context.Context, text string) (*models.Question, error) {
	return s.repository.Create(ctx, text)
}

func (s *QuestionService) GetAll(ctx context.Context) ([]*models.Question, error) {
	return s.repository.GetAll(ctx)
}

func NewQuestionService(repository repository.Question, service *Service) *QuestionService {
	return &QuestionService{
		service:    service,
		repository: repository,
	}
}
