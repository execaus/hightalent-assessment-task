package handler

import (
	"hightalent-assessment-task/internal/models"
	"hightalent-assessment-task/pkg/router"
)

func (h *Handler) CreateAnswer(ctx router.Context) {
	var request models.CreateAnswerRequest

	if err := ctx.BindJSON(&request); err != nil {
		ctx.Abort(err)
		return
	}

	userID, err := ctx.GetUserID()
	if err != nil {
		ctx.Abort(err)
		return
	}

	answer, err := h.service.Answer.Create(ctx, request.Text, request.QuestionID, userID)
	if err != nil {
		ctx.Abort(err)
		return
	}

	ctx.SendCreated(&models.CreateAnswerResponse{
		Answer: answer,
	})
}

func (h *Handler) GetAnswer(ctx router.Context) {
	answerID, err := ctx.GetIntDynamicValue("id")
	if err != nil {
		ctx.Abort(err)
		return
	}

	answer, err := h.service.Answer.Get(ctx, uint(answerID))
	if err != nil {
		if answer == nil {
			ctx.SendNotFound(err.Error())
			return
		}
		ctx.Abort(err)
		return
	}

	ctx.SendOK(&models.GetAnswerResponse{
		Answer: answer,
	})
}
