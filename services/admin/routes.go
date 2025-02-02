package admin

import (
	"fmt"
	"net/http"

	"github.com/XoliqberdiyevBehruz/wtc_backend/services/auth"
	types_admin "github.com/XoliqberdiyevBehruz/wtc_backend/types/admin"
	"github.com/XoliqberdiyevBehruz/wtc_backend/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
)

type Handler struct {
	store types_admin.UserStore
}

func NewHandler(store types_admin.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.HandleFunc("/user/create", auth.AuthWithJWT(h.handleCreateUser, h.store))
	r.HandleFunc("/user/login", h.handleLoginUser)
	r.HandleFunc("/user/profile", auth.AuthWithJWT(h.handleProfileUser, h.store))
	r.HandleFunc("/user/update/{userId}", auth.AuthWithJWT(h.handleUpdateUser, h.store))
	r.HandleFunc("/user/delete/{userId}", auth.AuthWithJWT(h.handleDeleteUser, h.store))
	r.HandleFunc("/user/detail/{userId}", auth.AuthWithJWT(h.handleDetailUser, h.store))
	r.HandleFunc("/user/list", auth.AuthWithJWT(h.handleGetAllUser, h.store))
}

// @Summary Create User
// @Description Create User
// @Tags admin-user
// @Accept  json
// @Produce  json
// @Param payload body types_admin.UserCreatePayload true "user create payload"
// @Router /user/create [post]
// @Security		BearerAuth
func (h *Handler) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var payload types_admin.UserCreatePayload

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload body: %v", errors))
		return
	}

	hashPassword, err := auth.GenerateHashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	u, err := h.store.GetUserByUsername(payload.Username)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error at get user by username: %v", err))
		return
	}

	if u != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user alreadt exists with this username"))
		return
	}

	user := types_admin.UserCreatePayload{
		Username:  payload.Username,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Password:  hashPassword,
	}

	err = h.store.CreateUser(&user)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to create user: %v", err))
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]string{"message": "user successfully created"})
}

// @Summary Login a user
// @Description Login a user with username and password
// @Tags admin-user
// @Accept  json
// @Produce  json
// @Param payload body types_admin.UserLoginPayload true "Login User Payload"
// @Router /user/login [post]
func (h *Handler) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	var payload types_admin.UserLoginPayload

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload bodyL %v", errors))
		return
	}

	user, err := h.store.GetUserByUsername(payload.Username)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user not found"))
		return
	}

	err = auth.CompareHashPassword(payload.Password, user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user not found"))
		return
	}

	token, err := auth.CreateJWT([]byte("jwt_token_secrets"), user.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"token": token})
}

// @Summary Get user profile
// @Tags admin-user
// @Accept  json
// @Produce  json
// @Router /user/profile [get]
// @Security		BearerAuth
func (h *Handler) handleProfileUser(w http.ResponseWriter, r *http.Request) {
	var userId = auth.GetUserIdFronContext(r.Context())

	user, err := h.store.GetUserById(userId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, user)
}

// @Summary Update user
// @Tags admin-user
// @Accept  json
// @Produce  json
// @Router /user/update/{userId} [put]
// @Param userId path string true "user Id"
// @Param payload body types_admin.UserUpdatePayload true "Update User Payload"
// @Security		BearerAuth
func (h *Handler) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	var payload types_admin.UserUpdatePayload
	var userId = r.PathValue("userId")

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalida payload body: %v", err))
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload body: %v", errors))
		return
	}

	user, err := h.store.GetUserById(userId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if user == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user not found"))
		return
	}

	userForUpdate := types_admin.UserUpdatePayload{
		Username:  payload.Username,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}

	err = h.store.UpdateUserById(&userForUpdate, user.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "successfully updated"})
}

// @Summary Delete user
// @Tags admin-user
// @Accept  json
// @Produce  json
// @Router /user/delete/{userId} [delete]
// @Param userId path string true "user Id"
// @Security		BearerAuth
func (h *Handler) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	var userId = r.PathValue("userId")

	user, err := h.store.GetUserById(userId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if user == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user not found"))
		return
	}

	err = h.store.DeleteUserById(userId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": "user successfully deleted"})
}

// @Summary Get user
// @Tags admin-user
// @Accept  json
// @Produce  json
// @Router /user/detail/{userId} [get]
// @Param userId path string true "user Id"
// @Security		BearerAuth
func (h *Handler) handleDetailUser(w http.ResponseWriter, r *http.Request) {
	var userId = r.PathValue("userId")

	user, err := h.store.GetUserById(userId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if user == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user not found"))
		return
	} else {
		utils.WriteJson(w, http.StatusOK, user)
	}
}

// @Summary List users
// @Tags admin-user
// @Accept  json
// @Produce  json
// @Router /user/list [get]
// @Security		BearerAuth
func (h *Handler) handleGetAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := h.store.GetAllUsers()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, users)
}
