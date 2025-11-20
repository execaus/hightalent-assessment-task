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

func (r *AnswerRepository) Delete(id uint) (*models.Answer, error) {
	ctx := context.Background()

	row, err := gorm.G[AnswerTable](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("answer not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get answer: %w", err)
	}

	_, err = gorm.G[AnswerTable](r.db).Where("id = ?", id).Delete(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to delete answer: %w", err)
	}

	return row.ToModel(), nil
}

func (r *AnswerRepository) GetAllByQuestionID(id uint) ([]*models.Answer, error) {
	rows, err := gorm.G[AnswerTable](r.db).Where("question_id = ?", id).Find(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get answers: %w", err)
	}

	answers := make([]*models.Answer, len(rows))
	for i := 0; i < len(rows); i++ {
		answers[i] = rows[i].ToModel()
	}

	return answers, nil
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
