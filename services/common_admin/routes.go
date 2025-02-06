package common_admin

import (
	"fmt"
	"io"
	"net/http"
	"os"

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
	// contactUsFooter
	r.Get("/admin/common/contact_us_footer/list", auth.AuthWithJWT(h.handleListContactUsFooter, h.userStore))
	r.Get("/admin/common/contact_us_footer/{contactId}", auth.AuthWithJWT(h.handleGetContactUsFooter, h.userStore))
	r.Put("/admin/common/contact_us_footer/{contactId}/update", auth.AuthWithJWT(h.handleUpdateContactUsFooter, h.userStore))
	r.Delete("/admin/common/contact_us_footer/{contactId}/delete", auth.AuthWithJWT(h.handleDeleteContactUsFooter, h.userStore))
	// medias
	r.Post("/admin/common/media/create", auth.AuthWithJWT(h.handleCreateMedia, h.userStore))
	r.Get("/admin/common/media/list", auth.AuthWithJWT(h.handleGetAllMedia, h.userStore))
	r.Delete("/admin/common/media/{mediaId}/delete", auth.AuthWithJWT(h.handleDeleteMedia, h.userStore))
	r.Put("/admin/common/media/{mediaId}/update", auth.AuthWithJWT(h.handleUpdateMedia, h.userStore))
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

// ========= ContactUsFooter ========= //
// @Summary ContactUsFooter list
// @Description ContactUsFooter list
// @Tags common-admin
// @Accept json
// @Produce json
// @Router /admin/common/contact_us_footer/list [get]
// @Security BearerAuth
func (h *Handler) handleListContactUsFooter(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.GetAllContactUsFooter()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, list)
}

// @Summary Get contactUs by id
// @Description Get contactUs by id
// @Tags common-admin
// @Accept json
// @Produce json
// @Router /admin/common/contact_us_footer/{contactId} [get]
// @Param contactId path string true "contact id"
// @Security BearerAuth
func (h *Handler) handleGetContactUsFooter(w http.ResponseWriter, r *http.Request) {
	var contactId = r.PathValue("contactId")

	contact, err := h.store.GetContactUsFooterById(contactId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, contact)
}

// @Summary update contact us
// @Description update contact us
// @Tags common-admin
// @Accept json
// @Produce json
// @Router /admin/common/contact_us_footer/{contactId}/update [put]
// @Param contactId path string true "contact Id"
// @Param payload body types_common_admin.ContactUsFooterUpdatePayload true "contact us update payload"
// @Security BearerAuth
func (h *Handler) handleUpdateContactUsFooter(w http.ResponseWriter, r *http.Request) {
	var payload types_common_admin.ContactUsFooterUpdatePayload
	var contactId = r.PathValue("contactId")

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	contactUs, err := h.store.GetContactUsFooterById(contactId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	changedContactUs := types_common_admin.ContactUsFooterUpdatePayload{
		IsContacted: payload.IsContacted,
	}

	err = h.store.UpdateContactUsFooter(contactUs.Id, &changedContactUs)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "successfully updated"})
}

// @Summary delete contact us
// @Description delete contact us
// @Tags common-admin
// @Accept json
// @Produce json
// @Router /admin/common/contact_us_footer/{contactId}/delete [delete]
// @Param contactId path string true "contact id"
// @Security BearerAuth
func (h *Handler) handleDeleteContactUsFooter(w http.ResponseWriter, r *http.Request) {
	var contactId = r.PathValue("contactId")
	contact, err := h.store.GetContactUsFooterById(contactId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.DeleteContactUsFooterById(contact.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": "successfully deleted"})
}

// @Summart create media
// @Description create media
// @Tags common-admin
// @Accept multipart/data
// @Produce json
// @Param fileUz formData file true "file uz"
// @Param fileRu formData file false "file ru"
// @Param fileEn formData file false "file en"
// @Router /admin/common/media/create [post]
// @Security BearerAuth
func (h *Handler) handleCreateMedia(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(50)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to parse form"))
		return
	}

	fileUz, fileHeader, err := r.FormFile("fileUz")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to read file_uz"))
		return
	}
	defer fileUz.Close()

	fileRu, fileRuHeader, err := r.FormFile("fileRu")
	if err != nil {
		if err == http.ErrMissingFile {
			fileRu = nil
			fileRuHeader = nil
		} else {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unabel to read file_ru"))
			return
		}
	}
	if fileRu != nil {
		defer fileRu.Close()
	}

	fileEn, fileEnHeader, err := r.FormFile("fileEn")
	if err != nil {
		if err == http.ErrMissingFile {
			fileEn = nil
			fileEnHeader = nil
		} else {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to read file_en"))
			return
		}
	}
	if fileEn != nil {
		defer fileEn.Close()
	}
	var fileRuPath string
	var fileEnPath string

	fileUzPath := "uploads/medias/files/" + fileHeader.Filename
	if fileRuHeader != nil {
		fileRuPath = "uploads/medias/files/" + fileRuHeader.Filename
	} else {
		fileRuPath = ""
	}
	if fileEnHeader != nil {
		fileEnPath = "uploads/medias/files/" + fileEnHeader.Filename
	} else {
		fileEnPath = ""
	}

	outFileUz, err := os.Create(fileUzPath)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to save fileUz: %v", err))
		return
	}
	defer outFileUz.Close()

	var outFileRu *os.File
	if fileRu != nil {
		outFileRu, err = os.Create(fileRuPath)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to save fileRu"))
			return
		}
		defer outFileRu.Close()
	}

	var outFileEn *os.File

	if fileEn != nil {

		outFileEn, err := os.Create(fileEnPath)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to save fileEn"))
			return
		}
		defer outFileEn.Close()
	}

	_, err = io.Copy(outFileUz, fileUz)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write fileUz"))
		return
	}

	if fileRu != nil {
		_, err = io.Copy(outFileRu, fileRu)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write fileRu"))
			return
		}
	}

	if fileEn != nil {
		_, err = io.Copy(outFileEn, fileEn)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write fileEn"))
			return
		}
	}

	meida := types_common_admin.MediaPayload{
		FileUz: fileUzPath,
		FileRu: fileRuPath,
		FileEn: fileEnPath,
	}

	err = h.store.CreateMedia(&meida)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

// @Summary get medias
// @Description get medias
// @Tags common-admin
// @Accept json
// @Produce json
// @Router /admin/common/media/list [get]
// @Security BearerAuth
func (h *Handler) handleGetAllMedia(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.GetAllMedias()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, list)
}

// @Summary delete media
// @Description delete media
// @Tags common-admin
// @Accept json
// @Produce json
// @Param mediaId path string true "media id"
// @Router /admin/common/media/{mediaId}/delete [delete]
// @Security BearerAuth
func (h *Handler) handleDeleteMedia(w http.ResponseWriter, r *http.Request) {
	var mediaId = r.PathValue("mediaId")

	media, err := h.store.GetMediaById(mediaId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.DeleteMediaById(media.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": "successfully deleted"})
}

// @Summary update media
// @Description update media
// @Tags common-admin
// @Accept multipart/data
// @Produce json
// @Param mediaId path string true "media id"
// @Param fileUz formData file true "file uz"
// @Param fileRu formData file false "file ru"
// @Param fileEn formData file false "file en"
// @Router /admin/common/media/{mediaId}/update [put]
// @Security BearerAuth
func (h *Handler) handleUpdateMedia(w http.ResponseWriter, r *http.Request) {
	var mediaId = r.PathValue("mediaId")
	media, err := h.store.GetMediaById(mediaId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = r.ParseMultipartForm(50)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	fileUz, fileHeader, err := r.FormFile("fileUz")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to read file_uz"))
		return
	}
	defer fileUz.Close()

	fileRu, fileRuHeader, err := r.FormFile("fileRu")
	if err != nil {
		if err == http.ErrMissingFile {
			fileRu = nil
			fileHeader = nil
		} else {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to read file_ru"))
			return
		}
	}
	if fileRu != nil {
		defer fileRu.Close()
	}

	fileEn, fileEnHeader, err := r.FormFile("fileEn")
	if err != nil {
		if err == http.ErrMissingFile {
			fileEn = nil
			fileEnHeader = nil
		} else {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to read file_en"))
			return
		}
	}

	if fileEn != nil {
		defer fileEn.Close()
	}

	var fileRuPath string
	var fileEnPath string

	fileUzPath := "uploads/medias/files/" + fileHeader.Filename
	if fileRuHeader != nil {
		fileRuPath = "uploads/medias/files/" + fileRuHeader.Filename
	} else {
		fileRuPath = ""
	}
	if fileEnHeader != nil {
		fileEnPath = "uploads/medias/files/" + fileEnHeader.Filename
	} else {
		fileEnPath = ""
	}

	outFileUz, err := os.Create(fileUzPath)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to save file_uz"))
		return
	}
	defer outFileUz.Close()

	var outFileRu *os.File
	if fileRuPath != "" {
		outFileRu, err = os.Create(fileRuPath)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to save file_ru"))
			return
		}
		defer outFileRu.Close()
	}

	var outFileEn *os.File
	if fileEnPath != "" {
		outFileEn, err = os.Create(fileEnPath)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to save file_en"))
			return
		}
		defer outFileEn.Close()
	}

	_, err = io.Copy(outFileUz, fileUz)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write file_uz"))
		return
	}

	if fileRuPath != "" {
		_, err = io.Copy(outFileRu, fileRu)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write file_ru"))
			return
		}
	}

	if fileEnPath != "" {
		_, err = io.Copy(outFileEn, fileEn)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write file_en"))
			return
		}
	}

	changed_media := types_common_admin.MediaPayload{
		FileUz: fileUzPath,
		FileRu: fileRuPath,
		FileEn: fileEnPath,
	}

	err = h.store.UpdateMedia(media.Id, &changed_media)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "successfully updated"})
}
