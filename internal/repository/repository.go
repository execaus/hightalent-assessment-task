package repository

import (
	"context"
	"hightalent-assessment-task/internal/models"
	"hightalent-assessment-task/pkg/router"

	"github.com/google/uuid"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Answer interface {
	Create(ctx context.Context, text string, questionID uint, userID uuid.UUID) (*models.Answer, error)
	Get(ctx context.Context, id uint) (*models.Answer, error)
	GetAllByQuestionID(id uint) ([]*models.Answer, error)
}

type User interface {
	IsExistByLogin(ctx context.Context, login string) (bool, error)
	Create(ctx context.Context, id uuid.UUID, login, password string) (*models.User, error)
	Get(ctx context.Context, id uuid.UUID) (*models.User, error)
}

type Question interface {
	IsExistByID(id uint) (bool, error)
	Create(ctx context.Context, text string) (*models.Question, error)
	GetAll(ctx context.Context) ([]*models.Question, error)
	Get(ctx router.Context, id uint) (*models.Question, error)
}

type Repository struct {
	Question
	User
	Answer
}
