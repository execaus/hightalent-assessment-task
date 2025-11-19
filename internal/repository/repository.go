package repository

import (
	"context"
	"hightalent-assessment-task/internal/models"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Question interface {
	Create(ctx context.Context, text string) (*models.Question, error)
}

type Repository struct {
	Question
}
