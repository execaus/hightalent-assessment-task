package handler

import (
	"hightalent-assessment-task/pkg/router"
	"strings"
)

const bearerPrefix = "Bearer "

func (h *Handler) requireAuth(ctx router.Context) {
	token, err := ctx.GetHeader("Authorization")
	if err != nil {
		ctx.Abort(router.NewUnauthorizedError(err.Error()))
		return
	}

	if !strings.HasPrefix(token, bearerPrefix) {
		ctx.Abort(router.NewUnauthorizedError("invalid Authorization header format"))
		return
	}

	token = strings.TrimPrefix(token, bearerPrefix)
	token = strings.TrimSpace(token)

	claims, err := h.service.Auth.GetClaims(token)
	if err != nil {
		ctx.Abort(router.NewUnauthorizedError(err.Error()))

		return
	}

	if err = ctx.SetUserID(claims.UserID); err != nil {
		ctx.Abort(router.NewUnauthorizedError(err.Error()))
		return
	}
}
