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

// TestCreateQuestion_HandlerCallServiceAndResponds проверяет, что handler CreateQuestion
// корректно вызывает метод сервиса Create и формирует правильный ответ
func TestCreateQuestion_HandlerCallServiceAndResponds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockQuestion(ctrl)

	services := service.Service{
		Question: mockService,
	}

	h := handler.NewHandler(&services)

	requestBody := models.CreateQuestionRequest{
		Text: "What is Go?",
	}

	requestContext := router.NewTestContext(t.Context(), time.Second)

	mockService.EXPECT().
		Create(requestContext, requestBody.Text).
		Return(&models.Question{
			ID:        0,
			Text:      "What is Go?",
			CreatedAt: time.Now(),
		}, nil)

	err := requestContext.PutRequestBody(requestBody)
	assert.NoError(t, err)

	h.CreateQuestion(requestContext)

	assert.False(t, requestContext.IsAbort())

	var responseBody models.CreateQuestionResponse

	err = json.Unmarshal(requestContext.ResponseBody, &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, responseBody.Question.ID, uint(0))
	assert.Equal(t, requestBody.Text, responseBody.Question.Text)
	assert.WithinDuration(t, time.Now(), responseBody.Question.CreatedAt, time.Second)
}

func TestGetAllQuestions_HandlerCallServiceAndResponds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockQuestion(ctrl)

	services := service.Service{
		Question: mockService,
	}

	h := handler.NewHandler(&services)

	expectedQuestions := []*models.Question{
		{ID: 1, Text: "What is Go?", CreatedAt: time.Now()},
		{ID: 2, Text: "What is Python?", CreatedAt: time.Now()},
	}

	requestContext := router.NewTestContext(t.Context(), time.Second)

	mockService.EXPECT().
		GetAll(requestContext).
		Return(expectedQuestions, nil)

	h.GetQuestions(requestContext)

	assert.False(t, requestContext.IsAbort())

	var responseBody models.GetAllQuestionsResponse
	err := json.Unmarshal(requestContext.ResponseBody, &responseBody)
	assert.NoError(t, err)

	assert.Len(t, responseBody.Questions, len(expectedQuestions))
	for i, q := range expectedQuestions {
		assert.Equal(t, q.ID, responseBody.Questions[i].ID)
		assert.Equal(t, q.Text, responseBody.Questions[i].Text)
		assert.WithinDuration(t, q.CreatedAt, responseBody.Questions[i].CreatedAt, time.Second)
	}
}
