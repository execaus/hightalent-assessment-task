package repository

import (
	"context"
	"errors"
	"fmt"
	"hightalent-assessment-task/internal/models"
	"hightalent-assessment-task/pkg/router"

	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

func (r *QuestionRepository) Delete(id uint) (*models.Question, error) {
	ctx := context.Background()

	row, err := gorm.G[QuestionTable](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("question not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get question: %w", err)
	}

	_, err = gorm.G[QuestionTable](r.db).Where("id = ?", id).Delete(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to delete question: %w", err)
	}

	return row.ToModel(), nil
}

func (r *QuestionRepository) Get(ctx router.Context, id uint) (*models.Question, error) {
	question, err := gorm.G[QuestionTable](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("question not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get question: %w", err)
	}
	return question.ToModel(), nil
}

func (r *QuestionRepository) IsExistByID(id uint) (bool, error) {
	var count int64
	if err := r.db.Model(&QuestionTable{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return false, fmt.Errorf("failed to check existence of question by id: %w", err)
	}
	return count > 0, nil
}

func (r *QuestionRepository) Create(ctx context.Context, text string) (*models.Question, error) {
	question := QuestionTable{
		Text: text,
	}

	if err := gorm.G[QuestionTable](r.db).Create(ctx, &question); err != nil {
		return nil, fmt.Errorf("failed to create question: %w", err)
	}

	return question.ToModel(), nil
}

func (r *QuestionRepository) GetAll(ctx context.Context) ([]*models.Question, error) {
	rows, err := gorm.G[QuestionTable](r.db).Find(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find questions: %w", err)
	}

	questions := make([]*models.Question, len(rows))
	for i := 0; i < len(rows); i++ {
		questions[i] = rows[i].ToModel()
	}

	return questions, nil
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{
		db: db,
	}
}
