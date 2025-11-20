package service

import (
	"context"
	"hightalent-assessment-task/internal/models"
	"hightalent-assessment-task/internal/repository"
)

type QuestionService struct {
	repository repository.Question
}

func (s *QuestionService) Create(ctx context.Context, text string) (*models.Question, error) {
	return s.repository.Create(ctx, text)
}

func (s *QuestionService) GetAll(ctx context.Context) ([]*models.Question, error) {
	return s.repository.GetAll(ctx)
}

func NewQuestionService(repository repository.Question) *QuestionService {
	return &QuestionService{
		repository: repository,
	}
}
