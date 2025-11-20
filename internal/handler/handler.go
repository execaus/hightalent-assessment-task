package handler

import (
	"hightalent-assessment-task/internal/service"
	"hightalent-assessment-task/pkg/router"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetRouter() *router.Router {
	r := router.DefaultRouter()

	api := r.Group("api")
	{
		v1 := api.Group("v1")
		{
			auth := v1.Group("auth")
			{
				auth.POST("sign-up", h.SignUp)
			}
			questions := v1.Group("questions")
			{
				questions.GET("", h.GetQuestions)
				questions.POST("", h.CreateQuestion)

				question := questions.Group("{id}")
				{
					question.GET("", h.GetQuestion)
					question.POST("answers", h.requireAuth, h.CreateAnswer)
				}
			}
			answers := v1.Group("answers")
			{
				answers.GET("{id}", h.GetAnswer)
			}
		}
	}

	return r
}
