package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"hightalent-assessment-task/internal/models"
	mockrepository "hightalent-assessment-task/internal/repository/mocks"
)

// TestQuestionService_Create проверяет, что сервис корректно вызывает репозиторий
// для создания вопроса и возвращает ожидаемый результат.
func TestQuestionService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuestion := mockrepository.NewMockQuestion(ctrl)

	svc := NewQuestionService(mockQuestion)

	text := "What is Go?"

	expected := &models.Question{
		Id:        1,
		Text:      text,
		CreatedAt: time.Now(),
	}

	mockQuestion.EXPECT().
		Create(text).
		Return(expected, nil).
		Times(1)

	result, err := svc.Create(text)

	assert.NoError(t, err)
	assert.Equal(t, expected.Id, result.Id)
	assert.Equal(t, expected.Text, result.Text)
	assert.WithinDuration(t, expected.CreatedAt, result.CreatedAt, time.Second)
}

// TestQuestionService_Create_Error проверяет, что сервис корректно возвращает
// ошибку, если репозиторий не смог создать вопрос.
func TestQuestionService_Create_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuestion := mockrepository.NewMockQuestion(ctrl)

	svc := NewQuestionService(mockQuestion)

	text := "bad"

	mockQuestion.EXPECT().
		Create(text).
		Return(nil, assert.AnError).
		Times(1)

	result, err := svc.Create(text)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
}
