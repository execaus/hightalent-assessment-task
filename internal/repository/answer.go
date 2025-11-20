package repository

import (
	"context"
	"errors"
	"fmt"
	"hightalent-assessment-task/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnswerRepository struct {
	db *gorm.DB
}

func (r *AnswerRepository) Get(ctx context.Context, id uint) (*models.Answer, error) {
	answer, err := gorm.G[AnswerTable](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("answer not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get answer: %w", err)
	}
	return answer.ToModel(), nil
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
