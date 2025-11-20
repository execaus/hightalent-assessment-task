package service

import (
	"context"
	"hightalent-assessment-task/internal/models"
	"hightalent-assessment-task/internal/repository"

	"github.com/google/uuid"
)

type AnswerService struct {
	service    *Service
	repository repository.Answer
}

func (s *AnswerService) Get(ctx context.Context, id uint) (*models.Answer, error) {
	return s.repository.Get(ctx, id)
}

func (s *AnswerService) Create(ctx context.Context, text string, questionID uint, userID uuid.UUID) (*models.Answer, error) {
	isExist, err := s.service.Question.IsExistByID(questionID)
	if err != nil {
		return nil, err
	}

	if !isExist {
		return nil, NewBusinessLoginError("cannot create answer: question does not exist")
	}

	answer, err := s.repository.Create(ctx, text, questionID, userID)
	if err != nil {
		return nil, err
	}

	return answer, err
}

func NewAnswerService(repository repository.Answer, service *Service) *AnswerService {
	return &AnswerService{
		service:    service,
		repository: repository,
	}
}
