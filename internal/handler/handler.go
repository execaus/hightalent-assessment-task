package handler

import "hightalent-assessment-task/internal/service"

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}
