package repository

import (
	"context"
	"hightalent-assessment-task/internal/models"

	"github.com/google/uuid"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type User interface {
	IsExistByLogin(ctx context.Context, login string) (bool, error)
	Create(ctx context.Context, id uuid.UUID, login, password string) (*models.User, error)
	Get(ctx context.Context, id uuid.UUID) (*models.User, error)
}

type Question interface {
	Create(ctx context.Context, text string) (*models.Question, error)
	GetAll(ctx context.Context) ([]*models.Question, error)
}

type Repository struct {
	Question
	User
}
