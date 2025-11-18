package service

import (
	"hightalent-assessment-task/internal/models"
	"hightalent-assessment-task/internal/repository"
)

type QuestionService struct {
	repository repository.Question
}

func (s *QuestionService) Create(test string) (*models.Question, error) {
	//TODO implement me
	panic("implement me")
}

func NewQuestionService(repository repository.Question) *QuestionService {
	return &QuestionService{
		repository: repository,
	}
}
