package repository

import (
	"context"
	"fmt"
	"hightalent-assessment-task/internal/models"

	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

func (r *QuestionRepository) Create(ctx context.Context, text string) (*models.Question, error) {
	question := QuestionTable{
		Text: text,
	}

	if err := gorm.G[QuestionTable](r.db).Create(ctx, &question); err != nil {
		return nil, fmt.Errorf("failed to create question: %w", err)
	}

	return &models.Question{
		ID:        question.ID,
		Text:      question.Text,
		CreatedAt: question.CreatedAt,
	}, nil
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{
		db: db,
	}
}
