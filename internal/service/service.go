package service

import (
	"context"
	"hightalent-assessment-task/config"
	"hightalent-assessment-task/internal/models"
	"hightalent-assessment-task/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Auth interface {
	GeneratePassword() (string, error)
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
	GenerateJWT(userID string) (string, error)
	GetClaims(tokenString string) (*models.AuthClaims, error)
}

type Question interface {
	Create(ctx context.Context, test string) (*models.Question, error)
	GetAll(ctx context.Context) ([]*models.Question, error)
}

type Service struct {
	Question
	Auth
}

func NewService(repos repository.Question, cfg *config.AuthConfig) *Service {
	return &Service{
		Question: NewQuestionService(repos),
		Auth:     NewAuthService(cfg),
	}
}
