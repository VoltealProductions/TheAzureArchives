package guild

import (
	"fmt"
	"log"
	"net/http"

	"github.com/VoltealProductions/TheAzureArcchives/middleware"
	"github.com/VoltealProductions/TheAzureArcchives/types"
	"github.com/VoltealProductions/TheAzureArcchives/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
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
		r.HandleFunc("GET /guild/show/{slug}", h.handleGetGuild)
		r.HandleFunc("POST /guild/create", h.handleCreateGuild)
	})
}

func (h *Handler) handleGetGuild(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	c, err := h.store.GetGuildBySlug(slug)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, c)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *Handler) handleCreateGuild(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateGuildPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.CreateGuild(types.Guild{
		OwnerId:     payload.OwnerId,
		Slug:        slug.Make(fmt.Sprintf("%s %s", payload.Name, payload.Realm)),
		Name:        payload.Name,
		Faction:     payload.Faction,
		Realm:       payload.Realm,
		Recruiting:  payload.Recruiting,
		Description: payload.Description,
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
