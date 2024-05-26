package delivery

import (
	"golang.org/x/exp/slog"

	"github.com/go-chi/chi"
)

type Handler struct {
	log *slog.Logger
}

func NewHandler(log *slog.Logger) *Handler {
	return &Handler{
		log: log,
	}
}

func (h *Handler) Handlers() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", h.Homepage)
	return r
}
