package user

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/VoltealProductions/TheAzureArcchives/config"
	"github.com/VoltealProductions/TheAzureArcchives/middleware"
	"github.com/VoltealProductions/TheAzureArcchives/services/auth"
	"github.com/VoltealProductions/TheAzureArcchives/types"
	"github.com/VoltealProductions/TheAzureArcchives/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.HandleFunc("POST /register", h.handleRegister)
	router.HandleFunc("POST /login", h.handleLogin)
	router.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.HandleFunc("PUT /user/update/{userId}", h.handleUpdateUser)
		r.HandleFunc("DELETE /user/delete/{userId}", h.handleDeleteUser)
		r.HandleFunc("POST /logout", h.handleLogin)
	})
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	u, err := h.store.GetUserByUsername(payload.Username)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if !auth.CompareHashedPasswords([]byte(u.Password), []byte(payload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("incorrect username or password"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, int(u.ID))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	err = utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
	if err != nil {
		log.Fatal(err)
	}
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	_, err := h.store.GetUserByUsername(payload.Username)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("a user that with username %s already exists", payload.Username))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	err = h.store.CreateUser(types.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: hashedPassword,
		Public:   payload.Public,
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

func (h *Handler) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	var payload types.UpdatePayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to parse user ID: %v", err))
		return
	}

	u, err := h.store.GetUserById(int(id))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("a user with id: %s does not exist", payload.Username))
		return
	}

	var pwd string
	if payload.Password == "" {
		pwd = u.Password
	} else {
		pwd, err = auth.HashPassword(payload.Password)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
		}
	}

	err = h.store.UpdateUser(int(id), types.User{
		Username: payload.Username,
		Password: pwd,
		Email:    payload.Email,
		Public:   payload.Public,
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

func (h *Handler) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to parse user ID: %v", err))
		return
	}

	err = h.store.DeleteUser(int(id))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, nil)
	if err != nil {
		log.Fatal(err)
	}
}
