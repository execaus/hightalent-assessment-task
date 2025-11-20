package service

import (
	"context"
	"hightalent-assessment-task/config"
	"hightalent-assessment-task/internal/models"
	"hightalent-assessment-task/internal/repository"
	"hightalent-assessment-task/pkg/router"

	"github.com/google/uuid"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Answer interface {
	Create(ctx context.Context, text string, questionID uint, userID uuid.UUID) (*models.Answer, error)
	Get(ctx context.Context, id uint) (*models.Answer, error)
	GetAllByQuestionID(id uint) ([]*models.Answer, error)
}

type User interface {
	IsExistByLogin(ctx context.Context, login string) (bool, error)
	Create(ctx context.Context, login, password string) (user *models.User, token string, err error)
	Get(ctx context.Context, id uuid.UUID) (*models.User, error)
}

type Auth interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
	GenerateJWT(userID string) (string, error)
	GetClaims(tokenString string) (*models.AuthClaims, error)
}

type Question interface {
	IsExistByID(id uint) (bool, error)
	Create(ctx context.Context, test string) (*models.Question, error)
	GetAll(ctx context.Context) ([]*models.Question, error)
	Get(ctx router.Context, id uint) (*models.Question, []*models.Answer, error)
}

type Service struct {
	Question
	Auth
	User
	Answer
}

func NewService(repos *repository.Repository, cfg *config.AuthConfig) *Service {
	s := Service{}

	s.Question = NewQuestionService(repos.Question, &s)
	s.Auth = NewAuthService(cfg)
	s.User = NewUserService(repos.User, &s)
	s.Answer = NewAnswerService(repos.Answer, &s)

	return &s
}
