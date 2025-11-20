package handler

import (
	"bytes"
	"encoding/json"
	"hightalent-assessment-task/internal/models"
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
	services := service.NewService(repositories, &cfg.Auth)
	handlers := NewHandler(services)

	router := handlers.GetRouter()

	body := models.CreateQuestionRequest{
		Text: "What is go?",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/questions", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.BaseHandle(w, req)

	require.Equal(t, http.StatusCreated, w.Code, "unexpected status code")

	var resp models.CreateQuestionResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err, "failed to unmarshal response")

	require.NotNil(t, resp.Question, "question should not be nil")
}

func TestGetAllQuestionsE2E(t *testing.T) {
	cfg := config.LoadTestConfig()

	repositories := repository.NewGormRepository(&cfg.Database)
	services := service.NewService(repositories, &cfg.Auth)
	handlers := NewHandler(services)

	router := handlers.GetRouter()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/questions", nil)
	w := httptest.NewRecorder()

	router.BaseHandle(w, req)

	require.Equal(t, http.StatusOK, w.Code, "unexpected status code")

	var resp models.GetAllQuestionsResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err, "failed to unmarshal response")

	require.NotNil(t, resp.Questions, "questions array should not be nil")
	require.GreaterOrEqual(t, len(resp.Questions), 1, "questions array should have at least one element")
}
