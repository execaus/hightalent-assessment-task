package repository

import (
	"fmt"
	"hightalent-assessment-task/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormRepository(cfg *config.DatabaseConfig) *Repository {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%v sslmode=disable",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.Port,
	)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	return &Repository{
		Question: NewQuestionRepository(gormDB),
		User:     NewUserRepository(gormDB),
	}
}
