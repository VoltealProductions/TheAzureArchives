package guild

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/VoltealProductions/TheAzureArcchives/middleware"
	"github.com/VoltealProductions/TheAzureArcchives/types"
	"github.com/VoltealProductions/TheAzureArcchives/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	slg "github.com/gosimple/slug"
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
		// r.HandleFunc("GET /guild/{slug}/members", h.handleGetGuild)
		r.HandleFunc("GET /user/{id}/guilds", h.HandleGetCharacterByUserId)
		r.HandleFunc("POST /guild/create", h.handleCreateGuild)
		r.HandleFunc("POST /guild/{slug}/update", h.handleUpdatGuild)
		r.HandleFunc("DELETE /guild/{slug}/delete", h.handleDeleteGuild)
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

func (h *Handler) HandleGetCharacterByUserId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	c, err := h.store.GetGuildsByUserId(uid)
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
		Slug:        slg.Make(fmt.Sprintf("%s %s", payload.Name, payload.Realm)),
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

func (h *Handler) handleUpdatGuild(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	exists, err := h.store.ConfirmThatGuildExists(slug)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if !exists {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("no guild with slug %s exists", slug))
		return
	}

	var payload types.UpdateGuildPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err = h.store.UpdateGuild(slug, types.Guild{
		Name:        payload.Name,
		Slug:        slg.Make(fmt.Sprintf("%s %s", payload.Name, payload.Realm)),
		Faction:     payload.Faction,
		Realm:       payload.Realm,
		Ranks:       payload.Ranks,
		Recruiting:  payload.Recruiting,
		Description: payload.Description,
		UpdatedAt:   time.Now(),
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

func (h *Handler) handleDeleteGuild(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	exists, err := h.store.ConfirmThatGuildExists(slug)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if !exists {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("no guild with slug %s exists", slug))
		return
	}

	err = h.store.DeleteGuild(slug)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, nil)
	if err != nil {
		log.Fatal(err)
	}
}
