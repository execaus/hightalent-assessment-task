package service

import (
	"context"
	"hightalent-assessment-task/internal/models"
	"hightalent-assessment-task/internal/repository"
	"hightalent-assessment-task/pkg/router"

	"github.com/google/uuid"
)

type UserService struct {
	service *Service
	repos   repository.User
}

func (s *UserService) IsExistByLogin(ctx context.Context, login string) (bool, error) {
	return s.repos.IsExistByLogin(ctx, login)
}

func (s *UserService) Create(ctx context.Context, login, password string) (*models.User, string, error) {
	isExist, err := s.service.User.IsExistByLogin(ctx, login)
	if err != nil {
		return nil, "", err
	}

	if isExist {
		return nil, "", router.NewBusinessLogicError("user with this login already exists")
	}

	passwordHash, err := s.service.Auth.HashPassword(password)
	if err != nil {
		return nil, "", err
	}

	user, err := s.repos.Create(ctx, uuid.New(), login, passwordHash)
	if err != nil {
		return nil, "", err
	}

	token, err := s.service.Auth.GenerateJWT(user.ID.String())
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *UserService) Get(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return s.repos.Get(ctx, id)
}

func NewUserService(repos repository.User, service *Service) *UserService {
	return &UserService{
		service: service,
		repos:   repos,
	}
}
