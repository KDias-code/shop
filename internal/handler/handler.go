package handler

import (
	"github.com/KDias-code/internal/service"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() http.Handler {
	router := chi.NewRouter()

	router.Route("/api", func(r chi.Router) {
		r.Post("/sign-up", h.SignUp)
		r.Post("/sign-in", h.SignIn)
		r.Post("/send-sms", h.Sms)
		r.Put("/{id}", h.Otps)
	})

	return router
}
