package service

import (
	"context"
	"hightalent-assessment-task/internal/models"
	"hightalent-assessment-task/internal/repository"
	"hightalent-assessment-task/pkg/router"
)

type QuestionService struct {
	service    *Service
	repository repository.Question
}

func (s *QuestionService) Delete(id uint) (*models.Question, error) {
	return s.repository.Delete(id)
}

func (s *QuestionService) Get(ctx router.Context, id uint) (*models.Question, []*models.Answer, error) {
	question, err := s.repository.Get(ctx, id)
	if err != nil {
		return nil, nil, err
	}

	answers, err := s.service.Answer.GetAllByQuestionID(id)

	return question, answers, err
}

func (s *QuestionService) IsExistByID(id uint) (bool, error) {
	return s.repository.IsExistByID(id)
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
