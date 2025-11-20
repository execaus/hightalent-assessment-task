package repository

import (
	"context"
	"errors"
	"fmt"
	"hightalent-assessment-task/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) IsExistByLogin(ctx context.Context, login string) (bool, error) {
	var count int64

	err := r.db.WithContext(ctx).Model(&UserTable{}).
		Where("login = ?", login).
		Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("failed to check if user exists: %w", err)
	}

	return count > 0, nil
}

func (r *UserRepository) Create(ctx context.Context, id uuid.UUID, login, password string) (*models.User, error) {
	user := UserTable{
		ID:       id,
		Login:    login,
		Password: password,
	}

	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user.ToModel(), nil
}

func (r *UserRepository) Get(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user UserTable
	if err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user.ToModel(), nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
