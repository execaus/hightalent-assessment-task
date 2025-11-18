package repository

import "hightalent-assessment-task/internal/models"

type QuestionRepository struct {
}

func (r *QuestionRepository) Create(text string) (*models.Question, error) {
	//TODO implement me
	panic("implement me")
}

func NewQuestionRepository() *QuestionRepository {
	return &QuestionRepository{}
}
