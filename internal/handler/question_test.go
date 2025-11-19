package handler_test

import (
	"encoding/json"
	"hightalent-assessment-task/internal/handler"
	"hightalent-assessment-task/internal/models"
	"hightalent-assessment-task/internal/service"
	mock_service "hightalent-assessment-task/internal/service/mocks"
	"hightalent-assessment-task/pkg/router"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// TestCreateQuestion_HandlerCallsServiceAndResponds проверяет, что handler CreateQuestion
// корректно вызывает метод сервиса Create и формирует правильный ответ
func TestCreateQuestion_HandlerCallServiceAndResponds(t *testing.T) {
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
			ID:        0,
			Text:      "What is Go?",
			CreatedAt: time.Now(),
		}, nil)

	requestContext := router.NewTestContext(t.Context())

	err := requestContext.PutRequestBody(requestBody)
	assert.NoError(t, err)

	h.CreateQuestion(requestContext)

	assert.False(t, requestContext.IsAbort())

	var responseBody models.CreateQuestionResponse

	err = json.Unmarshal(requestContext.Response.Body, &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, responseBody.Question.ID, 0)
	assert.Equal(t, requestBody.Text, responseBody.Question.Text)
	assert.WithinDuration(t, time.Now(), responseBody.Question.CreatedAt, time.Second)
}
