package main

import (
	"context"
	"embed"
	"errors"
	"hightalent-assessment-task/config"
	"hightalent-assessment-task/internal/handler"
	"hightalent-assessment-task/internal/repository"
	"hightalent-assessment-task/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	cfg := config.LoadConfig()

	repositories, gormDB := repository.NewGormRepository(&cfg.Database)

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}
	db, err := gormDB.DB()
	if err != nil {
		log.Fatal(err)
	}
	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal(err)
	}

	services := service.NewService(repositories, &cfg.Auth)
	handlers := handler.NewHandler(services)

	r := handlers.GetRouter()
	r.PrintRoutes()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	server := r.GetServer(&cfg.Server.Port)

	go func() {
		if err = server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not listen on %s: %v\n", cfg.Server.Port, err)
		}
	}()

	<-stop
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server gracefully stopped")
}
