package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"hightalent-assessment-task/config"
	"hightalent-assessment-task/internal/repository"
	"hightalent-assessment-task/internal/service"
)

func TestCreateQuestionE2E(t *testing.T) {
	cfg := config.LoadTestConfig()

	repositories := repository.NewGormRepository(&cfg.Database)
	services := service.NewService(repositories)
	handlers := NewHandler(services)

	router := handlers.GetRouter()

	body := map[string]interface{}{
		"test": "What is go?",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/questions", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.BaseHandle(w, req)

	require.Equal(t, http.StatusCreated, w.Code, "unexpected status code")
	require.NotEmpty(t, w.Body.String(), "response body should not be empty")
}
