package common

import (
	"fmt"
	"net/http"

	types_common "github.com/XoliqberdiyevBehruz/wtc_backend/types/common"
	"github.com/XoliqberdiyevBehruz/wtc_backend/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
)

type Handler struct {
	store types_common.CommonStore
}

func NewHandler(store types_common.CommonStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Post("/contact_us/create", h.handleContactUsCreate)
	r.Get("/settings", h.handleGetSettings)
	r.Post("/contact_us_footer/create", h.handleContactUsFooterCreate)
	r.Get("/media/list", h.handleGetMediasList)
	r.Get("/partner/list", h.handleListPartner)
}

// @Summary create contact us
// @Description create contact us
// @Tags common
// @Accept  json
// @Produce  json
// @Param payload body types_common.ContactCreatePayload true "Contact Us Payload"
// @Router /contact_us/create [post]
func (h *Handler) handleContactUsCreate(w http.ResponseWriter, r *http.Request) {
	var payload types_common.ContactCreatePayload

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload body: %v", errors))
		return
	}

	contact := types_common.ContactCreatePayload{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Comment:   payload.Comment,
	}

	err := h.store.CreateContactUs(&contact)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

// <==SETTINGS==> //
// @Summary get settings
// @Description get settings list
// @Tags common
// @Accept json
// @Produce json
// @Router /settings [get]
func (h *Handler) handleGetSettings(w http.ResponseWriter, r *http.Request) {
	settings, err := h.store.GetAllSettings()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, settings)
}

// @Summary create contact footer
// @Description create contact footer
// @Tags common
// @Accept json
// @Produce json
// @Param payload body types_common.ContactUsFooterPayload true "contact us footer create payload"
// @Router /contact_us_footer/create [post]
func (h *Handler) handleContactUsFooterCreate(w http.ResponseWriter, r *http.Request) {
	var payload types_common.ContactUsFooterPayload

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	contactUs := types_common.ContactUsFooterPayload{
		FullName: payload.FullName,
		Phone:    payload.Phone,
		Email:    payload.Email,
	}

	err := h.store.CreateContactUsFooter(contactUs)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

// @Summary get all media
// @Description get all media
// @Tags common
// @Accept json
// @Produce json
// @Router /media/list [get]
func (h *Handler) handleGetMediasList(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.GetAllMedia()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, list)
}

// @Summary list partner
// @Description list partner
// @Tags common
// @Accept json
// @Produce json
// @Router /partner/list [get]
func (h *Handler) handleListPartner(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.ListPartner()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, list)
}
