package common_admin

import (
	"fmt"
	"net/http"

	"github.com/XoliqberdiyevBehruz/wtc_backend/services/auth"
	types_common_admin "github.com/XoliqberdiyevBehruz/wtc_backend/types/common_admin"
	types_admin "github.com/XoliqberdiyevBehruz/wtc_backend/types/user_admin"
	"github.com/XoliqberdiyevBehruz/wtc_backend/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
)

type Handler struct {
	store     types_common_admin.CommonStore
	userStore types_admin.UserStore
}

func NewHandler(store types_common_admin.CommonStore, userStore types_admin.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	// contact us
	r.Get("/admin/common/contact_us/list", auth.AuthWithJWT(h.handleGetContactUsList, h.userStore))
	r.Delete("/admin/common/contact_us/{contactId}/delete", auth.AuthWithJWT(h.handleDeleteContactUs, h.userStore))
	r.Put("/admin/common/contact_us/{contactId}/update", auth.AuthWithJWT(h.handleUpdateContactUs, h.userStore))
	r.Get("/admin/common/contact_us/{contactId}", auth.AuthWithJWT(h.handleGetContactUs, h.userStore))
	// settings
	r.Post("/admin/common/settings/create", auth.AuthWithJWT(h.handleCreateSettings, h.userStore))
	r.Get("/admin/common/settings/{settingsId}", auth.AuthWithJWT(h.handleGetSettings, h.userStore))
	r.Put("/admin/common/settings/{settingsId}/update", auth.AuthWithJWT(h.handleUpdateSettings, h.userStore))
	r.Get("/admin/common/settings/list", auth.AuthWithJWT(h.handleListSettings, h.userStore))
}

// =========================== Contact Us ===========================

// @Summary get contact us
// @Description get contact us
// @Tags common-admin
// @Accept  json
// @Produce  json
// @Router /admin/common/contact_us/list [get]
// @Security		BearerAuth
func (h *Handler) handleGetContactUsList(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.GetContactUsList()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, list)
}

// @Summary delete contact us
// @Description delete contact us
// @Tags common-admin
// @Accept  json
// @Produce  json
// @Router /admin/common/contact_us/{contactId}/delete [delete]
// @Param contactId path string true "contact Id"
// @Security		BearerAuth
func (h *Handler) handleDeleteContactUs(w http.ResponseWriter, r *http.Request) {
	var contactId = r.PathValue("contactId")

	contact, err := h.store.GetContactUsById(contactId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.DeleteContactUsById(contact.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": "successfully deleted"})
}

// @Summary update contact us
// @Description update contact us
// @Tags common-admin
// @Accept  json
// @Produce  json
// @Router /admin/common/contact_us/{contactId}/update [put]
// @Param contactId path string true "contact Id"
// @Param payload body types_common_admin.ContactUpdatePayload true "Update Contact Payload"
// @Security		BearerAuth
func (h *Handler) handleUpdateContactUs(w http.ResponseWriter, r *http.Request) {
	var payload types_common_admin.ContactUpdatePayload
	var contactId = r.PathValue("contactId")
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid paylaod body: %v", errors))
		return
	}

	contact, err := h.store.GetContactUsById(contactId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.UpdateContactById(contact.Id, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "successfully updated"})
}

// @Summary get contact us
// @Description get contact us
// @Tags common-admin
// @Accept  json
// @Produce  json
// @Router /admin/common/contact_us/{contactId} [get]
// @Param contactId path string true "contact Id"
// @Security		BearerAuth
func (h *Handler) handleGetContactUs(w http.ResponseWriter, r *http.Request) {
	var contactId = r.PathValue("contactId")

	contact, err := h.store.GetContactUsById(contactId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, contact)
}

// =========================== Settings ===========================
// @Summary create settings
// @Description create settings
// @Tags common-admin
// @Accept  json
// @Produce  json
// @Router /admin/common/settings/create [post]
// @Param payload body types_common_admin.SettingsCreatePayload true "Create Settings Payload"
// @Security		BearerAuth
func (h *Handler) handleCreateSettings(w http.ResponseWriter, r *http.Request) {
	var payload types_common_admin.SettingsCreatePayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload body: %v", errors))
		return
	}

	settings := types_common_admin.SettingsCreatePayload{
		FirstPhone:  payload.FirstPhone,
		SecondPhone: payload.SecondPhone,
		Email:       payload.Email,
		Telegram:    payload.Telegram,
		Instagram:   payload.Instagram,
		Youtube:     payload.Youtube,
		Facebook:    payload.Facebook,
		AddressUz:   payload.AddressUz,
		AddressRu:   payload.AddressRu,
		AddressEn:   payload.AddressEn,
		WorkingDays: payload.WorkingDays,
	}
	err := h.store.CreateSettings(&settings)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]string{"message": "succesfully created"})
}

// @Summary get settings
// @Description get settings
// @Tags common-admin
// @Accept  json
// @Produce  json
// @Router /admin/common/settings/list [get]
// @Security		BearerAuth
func (h *Handler) handleListSettings(w http.ResponseWriter, r *http.Request) {
	settings, err := h.store.GetSettings()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, settings)
}

// @Summary settings update
// @Description settings update by id
// @Tags common-admin
// @Accept json
// @Produce json
// @Router /admin/common/settings/{settingsId}/update [put]
// @Param settingsId path string true "settings Id"
// @Param payload body types_common_admin.SettingsUpdatePayload true "settings update payload"
// @Security BearerAuth
func (h *Handler) handleUpdateSettings(w http.ResponseWriter, r *http.Request) {
	var payload types_common_admin.SettingsUpdatePayload
	var settingsId = r.PathValue("settingsId")

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, errors)
		return
	}

	settings, err := h.store.GetSettingsById(settingsId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	changed_settings := types_common_admin.SettingsUpdatePayload{
		FirstPhone:  payload.FirstPhone,
		SecondPhone: payload.SecondPhone,
		Email:       payload.Email,
		Telegram:    payload.Telegram,
		Instagram:   payload.Instagram,
		Youtube:     payload.Youtube,
		Facebook:    payload.Facebook,
		AddressUz:   payload.AddressUz,
		AddressRu:   payload.AddressRu,
		AddressEn:   payload.AddressEn,
		WorkingDays: payload.WorkingDays,
	}

	err = h.store.UpdateSettingsById(settings.Id, &changed_settings)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "successfully updated"})
}

// @Summary settings get
// @Description settings get with setting id
// @Tags common-admin
// @Accept json
// @Produce json
// @Router /admin/common/settings/{settingsId} [get]
// @Param settingsId path string true "settings Id"
// @Security BearerAuth
func (h *Handler) handleGetSettings(w http.ResponseWriter, r *http.Request) {
	var settingsId = r.PathValue("settingsId")
	settings, err := h.store.GetSettingsById(settingsId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, settings)
}
