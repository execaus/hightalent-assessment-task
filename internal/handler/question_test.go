package handler_test

import (
	"hightalent-assessment-task/internal/handler"
	"hightalent-assessment-task/internal/models"
	"hightalent-assessment-task/internal/service"
	mock_service "hightalent-assessment-task/internal/service/mocks"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateQuestion_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockQuestion(ctrl)

	services := service.Service{
		Question: mockService,
	}

	h := handler.NewHandler(services)

	requestBody := models.CreateQuestionRequest{
		Text: "What is Go?",
	}

	mockService.EXPECT().
		Create(gomock.Any()).
		Return(int64(1), nil)

	// TODO response := h.CreateQuestion()

	assert.Equal(t, http.StatusCreated, response.StatusCode)

	assert.NoError(t, err)
	assert.Greater(t, response.ID, int64(0))
	assert.Equal(t, requestBody.Text, response.Body.Text)
	assert.WithinDuration(t, time.Now(), response.Body.CreatedAt, time.Second)
}
