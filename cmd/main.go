package main

import (
	"context"
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
)

func main() {
	cfg := config.LoadConfig()

	repositories := repository.NewGormRepository(&cfg.Database)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	r := handlers.GetRouter()
	r.PrintRoutes()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	server := r.GetServer(&cfg.Server.Port)

	go func() {
		log.Printf("Server is running on port %s\n", cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
