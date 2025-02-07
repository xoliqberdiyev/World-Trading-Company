package about_company

import (
	"fmt"
	"net/http"

	"github.com/XoliqberdiyevBehruz/wtc_backend/services/auth"
	types_about_company "github.com/XoliqberdiyevBehruz/wtc_backend/types/about_company"
	types_user_admin "github.com/XoliqberdiyevBehruz/wtc_backend/types/user_admin"
	"github.com/XoliqberdiyevBehruz/wtc_backend/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
)

type Handler struct {
	store     types_about_company.CompanyStore
	userStore types_user_admin.UserStore
}

func NewHandler(store types_about_company.CompanyStore, userStore types_user_admin.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Post("/admin/company/capasity/create", auth.AuthWithJWT(h.handleCreateCapasity, h.userStore))
	r.Get("/admin/company/capasity/list", auth.AuthWithJWT(h.handleListCapasity, h.userStore))
	r.Delete("/admin/company/capasity/{capasityId}/delete", auth.AuthWithJWT(h.handleDeleteCapasity, h.userStore))
	r.Put("/admin/company/capasity/{capasityId}/update", auth.AuthWithJWT(h.handleUpdateCapasity, h.userStore))
	r.Get("/admin/company/capasity/{capasityId}", auth.AuthWithJWT(h.handleGetCapasity, h.userStore))
}

// @Summary create capasity
// @Description create capasity
// @Tags company-admin
// @Accept json
// @Produce json
// @Router /admin/company/capasity/create [post]
// @Param payload body types_about_company.CapasityPayload true "payload"
// @Security BearerAuth
func (h *Handler) handleCreateCapasity(w http.ResponseWriter, r *http.Request) {
	var payload types_about_company.CapasityPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	capasity := types_about_company.CapasityPayload{
		NameUz:   payload.NameUz,
		NameRu:   payload.NameRu,
		NameEn:   payload.NameEn,
		Quantity: payload.Quantity,
	}

	err := h.store.CreateCapasity(&capasity)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

// @Summary list capasity
// @Description list capasity
// @Tags company-admin
// @Accept json
// @Produce json
// @Router /admin/company/capasity/list [get]
// @Security BearerAuth
func (h *Handler) handleListCapasity(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.ListCapasity()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, list)
}

// @Summary delete capasity
// @Description delete capasity
// @Tags company-admin
// @Accept json
// @Produce json
// @Param capasityId path string true "capasity id"
// @Router /admin/company/capasity/{capasityId}/delete [delete]
// @Security BearerAuth
func (h *Handler) handleDeleteCapasity(w http.ResponseWriter, r *http.Request) {
	var capasityId = r.PathValue("capasityId")
	capasity, err := h.store.GetCapasity(capasityId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if capasity == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("capasity not found"))
		return
	}

	err = h.store.DeleteCapasity(capasity.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": "successfully deleted"})
}

// @Summary update capasity
// @Description update capasity
// @Tags company-admin
// @Accept json
// @Produce json
// @Param capasityId path string true "capasity id"
// @Param payload body types_about_company.CapasityPayload true "update capasity payload"
// @Router /admin/company/capasity/{capasityId}/update [put]
// @Security BearerAuth
func (h *Handler) handleUpdateCapasity(w http.ResponseWriter, r *http.Request) {
	var capasityId = r.PathValue("capasityId")
	var payload types_about_company.CapasityPayload
	capasity, err := h.store.GetCapasity(capasityId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if capasity == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("capasity not found"))
		return
	}

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	changed_capasity := types_about_company.CapasityPayload{
		NameUz:   payload.NameUz,
		NameRu:   payload.NameRu,
		NameEn:   payload.NameEn,
		Quantity: payload.Quantity,
	}
	err = h.store.UpdateCapasity(capasity.Id, &changed_capasity)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "successfully updated"})
}

// @Summary get capasity
// @Description get capasity
// @Tags company-admin
// @Accept json
// @Produce json
// @Param capasityId path string true "capasity id"
// @Router /admin/company/capasity/{capasityId} [get]
// @Security BearerAuth
func (h *Handler) handleGetCapasity(w http.ResponseWriter, r *http.Request) {
	var capasityId = r.PathValue("capasityId")
	capasity, err := h.store.GetCapasity(capasityId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, capasity)
}
