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
			questions := v1.Group("questions")
			{
				questions.POST("", h.CreateQuestion)
			}
		}
	}

	return r
}
