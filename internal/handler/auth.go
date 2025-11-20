package handler

import (
	"hightalent-assessment-task/internal/models"
	"hightalent-assessment-task/pkg/router"
)

func (h *Handler) SignUp(ctx router.Context) {
	var request models.SignUpRequest

	if err := ctx.BindJSON(&request); err != nil {
		ctx.Abort(err)
		return
	}

	user, token, err := h.service.User.Create(ctx, request.Login, request.Password)
	if err != nil {
		ctx.Abort(err)
		return
	}

	ctx.SendCreated(&models.SignUpResponse{
		User:  user,
		Token: token,
	})
}
