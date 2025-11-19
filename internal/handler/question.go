package handler

import (
	"hightalent-assessment-task/internal/models"
	"hightalent-assessment-task/pkg/router"
)

func (h *Handler) CreateQuestion(ctx router.Context) {
	var request models.CreateQuestionRequest

	if err := ctx.BindJSON(&request); err != nil {
		ctx.Abort(err)
		return
	}

	question, err := h.service.Question.Create(request.Text)
	if err != nil {
		ctx.Abort(err)
		return
	}

	ctx.SendCreated(&models.CreateQuestionResponse{
		Question: question,
	})
}
