package repository

import (
	"context"
	"fmt"
	"hightalent-assessment-task/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnswerRepository struct {
	db *gorm.DB
}

func (r *AnswerRepository) Create(ctx context.Context, text string, questionID uint, userID uuid.UUID) (*models.Answer, error) {
	answer := AnswerTable{
		Text:       text,
		QuestionID: questionID,
		UserID:     userID,
	}

	if err := gorm.G[AnswerTable](r.db).Create(ctx, &answer); err != nil {
		return nil, fmt.Errorf("failed to create answer: %w", err)
	}

	return answer.ToModel(), nil
}

func NewAnswerRepository(db *gorm.DB) *AnswerRepository {
	return &AnswerRepository{
		db: db,
	}
}
