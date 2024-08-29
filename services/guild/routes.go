package guild

import (
	"github.com/VoltealProductions/TheAzureArcchives/middleware"
	"github.com/VoltealProductions/TheAzureArcchives/types"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store types.GuildStore
}

func NewHandler(store types.GuildStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		// r.HandleFunc("POST /create/character", h.handleCreateCharacter)
	})
}
