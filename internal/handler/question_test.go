package handler_test

import (
	"encoding/json"
	"hightalent-assessment-task/internal/handler"
	"hightalent-assessment-task/internal/models"
	"hightalent-assessment-task/internal/service"
	mock_service "hightalent-assessment-task/internal/service/mocks"
	"hightalent-assessment-task/pkg/router"
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
		Create(requestBody.Text).
		Return(&models.Question{
			ID:        1,
			Text:      "What is Go?",
			CreatedAt: time.Now(),
		}, nil)

	requestContext := router.NewTestContext(t.Context())

	h.CreateQuestion(requestContext)

	assert.False(t, requestContext.IsAbort())

	assert.Equal(t, http.StatusCreated, requestContext.Response.StatusCode)

	var responseBody models.CreateQuestionResponse
	if err := json.Unmarshal(requestContext.Response.Body, &responseBody); err != nil {
		assert.NoError(t, err)
	}

	assert.Greater(t, responseBody.Question.ID, int64(0))
	assert.Equal(t, requestBody.Text, responseBody.Question.Text)
	assert.WithinDuration(t, time.Now(), responseBody.Question.CreatedAt, time.Second)
}
