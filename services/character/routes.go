package character

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
		r.HandleFunc("GET /character/show/{id}", h.HandleGetCharacter)
		r.HandleFunc("GET /user/{id}/characters", h.HandleGetCharacterByUserId)
		r.HandleFunc("PUT /character/update/{id}", h.handleUpdateCharacter)
		r.HandleFunc("DELETE /character/delete/{id}", h.handleDeleteCharacter)
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

	err = utils.WriteJSON(w, http.StatusCreated, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *Handler) HandleGetCharacter(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	c, err := h.store.GetCharacterByUniqueId(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, c)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *Handler) HandleGetCharacterByUserId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	c, err := h.store.GetCharacterByUserId(uid)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, c)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *Handler) handleUpdateCharacter(w http.ResponseWriter, r *http.Request) {
	charID := chi.URLParam(r, "id")
	var payload types.UpdateCharacterPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	id, err := strconv.ParseInt(charID, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to parse user ID: %v", err))
		return
	}

	err = h.store.UpdateCharacter(int(id), types.Character{
		Firstname:  payload.Firstname,
		Lastname:   payload.Lastname,
		Faction:    payload.Faction,
		Class:      payload.Class,
		Species:    payload.Species,
		ShortTitle: payload.ShortTitle,
		FullTitle:  payload.FullTitle,
		Age:        payload.Age,
		Gender:     payload.Gender,
		Pronouns:   payload.Pronouns,
		Height:     payload.Height,
		Weight:     payload.Weight,
		Birthplace: payload.Birthplace,
		Residence:  payload.Residence,
		About:      payload.About,
		History:    payload.History,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *Handler) handleDeleteCharacter(w http.ResponseWriter, r *http.Request) {
	charID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(charID, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to parse user ID: %v", err))
		return
	}

	err = h.store.DeleteCharacter(int(id))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, nil)
	if err != nil {
		log.Fatal(err)
	}
}
