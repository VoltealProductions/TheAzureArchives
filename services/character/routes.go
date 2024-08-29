package character

import (
	"fmt"
	"net/http"

	"github.com/VoltealProductions/TheAzureArcchives/middleware"
	"github.com/VoltealProductions/TheAzureArcchives/types"
	"github.com/VoltealProductions/TheAzureArcchives/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
)

type Handler struct {
	store types.CharacterStore
}

func NewHandler(store types.CharacterStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.HandleFunc("POST /create/character", h.handleCreateCharacter)
	})
}

func (h *Handler) handleCreateCharacter(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateCharacterPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.CreateCharacter(types.Character{
		UserID:    1,
		UniqueId:  utils.GenerateURLId(),
		Firstname: payload.Firstname,
		Lastname:  payload.Lastname,
		Faction:   payload.Faction,
		Species:   payload.Species,
		Class:     payload.Class,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
