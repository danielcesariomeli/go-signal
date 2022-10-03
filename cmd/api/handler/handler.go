package handler

import "net/http"

type Service interface {
	StartCounter() error
	StopCounter() error
}

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) HandleStart(w http.ResponseWriter, r *http.Request) {
	h.Service.StartCounter()
}

func (h *Handler) HandleStop(w http.ResponseWriter, r *http.Request) {
	h.Service.StopCounter()
}
