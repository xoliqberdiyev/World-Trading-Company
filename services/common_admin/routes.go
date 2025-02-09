package common_admin

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

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
	r.Get("/admin/common/media/{mediaId}", auth.AuthWithJWT(h.handleGetMedia, h.userStore))
	//partner
	r.Post("/admin/common/partner/create", auth.AuthWithJWT(h.handleCreatePartner, h.userStore))
	r.Get("/admin/common/partner/list", auth.AuthWithJWT(h.handleListPartner, h.userStore))
	r.Put("/admin/common/partner/{partnerId}/update", auth.AuthWithJWT(h.handleUpdatePartner, h.userStore))
	r.Delete("/admin/common/partner/{partnerId}/delete", auth.AuthWithJWT(h.handleDeletePartner, h.userStore))
	r.Get("/admin/common/partner/{partnerId}", auth.AuthWithJWT(h.handleGetPartner, h.userStore))
	// banner
	r.Post("/admin/common/banner/create", auth.AuthWithJWT(h.handleCreateBanner, h.userStore))
	r.Get("/admin/common/banner/list", auth.AuthWithJWT(h.handleListBanner, h.userStore))
	r.Put("/admin/common/banner/{bannerId}/update", auth.AuthWithJWT(h.handleUpdateBanner, h.userStore))
	r.Delete("/admin/common/banner/{bannerId}/delete", auth.AuthWithJWT(h.handleDeleteBanner, h.userStore))
	r.Get("/admin/common/banner/{bannerId}", auth.AuthWithJWT(h.handleGetBanner, h.userStore))
	// news
	r.Post("/admin/common/news/create", auth.AuthWithJWT(h.handleCreateNews, h.userStore))
	r.Get("/admin/common/news/list", auth.AuthWithJWT(h.handleListNews, h.userStore))
	r.Put("/admin/common/news/{newsId}/update", auth.AuthWithJWT(h.handleUpdateNews, h.userStore))
	r.Delete("/admin/common/news/{newsId}/delete", auth.AuthWithJWT(h.handleDeleteNews, h.userStore))
	r.Get("/admin/common/news/{newsId}", auth.AuthWithJWT(h.handleGetNews, h.userStore))
	// certificate
	r.Post("/admin/common/certificate/create", auth.AuthWithJWT(h.handleCreateCertificate, h.userStore))
	r.Get("/admin/common/certificate/list", auth.AuthWithJWT(h.handleListCertificate, h.userStore))
	r.Delete("/admin/common/certificate/{certificateId}/delete", auth.AuthWithJWT(h.handleDeleteCertificate, h.userStore))
	r.Put("/admin/common/certificate/{certificateId}/update", auth.AuthWithJWT(h.handleUpdateCertificate, h.userStore))
	r.Get("/admin/common/certificate/{certificateId}", auth.AuthWithJWT(h.handleGetCertificate, h.userStore))
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
// @Param link formData string false "link"
// @Router /admin/common/media/create [post]
// @Security BearerAuth
func (h *Handler) handleCreateMedia(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(50)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to parse form"))
		return
	}
	link := r.FormValue("link")
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
		outFileEn, err = os.Create(fileEnPath)
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
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write fileEn: %v", err))
			return
		}
	}

	meida := types_common_admin.MediaPayload{
		FileUz: fileUzPath,
		FileRu: fileRuPath,
		FileEn: fileEnPath,
		Link:   link,
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
// @Param link formData string false "link"
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
	link := r.FormValue("link")

	changed_media := types_common_admin.MediaPayload{
		FileUz: fileUzPath,
		FileRu: fileRuPath,
		FileEn: fileEnPath,
		Link:   link,
	}

	err = h.store.UpdateMedia(media.Id, &changed_media)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "successfully updated"})
}

// @Summary get media
// @Description get media
// @Tags common-admin
// @Accept json
// @Produce json
// @Param mediaId path string true "media id"
// @Router /admin/common/media/{mediaId} [get]
// @Security BearerAuth
func (h *Handler) handleGetMedia(w http.ResponseWriter, r *http.Request) {
	var mediaId = r.PathValue("mediaId")
	media, err := h.store.GetMediaById(mediaId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, media)
}

// @Summary create partner
// @Description create partner
// @Tags common-admin
// @Accept multipart/data
// @Produce json
// @Param image formData file true "image"
// @Router /admin/common/partner/create [post]
// @Security BearerAuth
func (h *Handler) handleCreatePartner(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(50)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	image, imageHeader, err := r.FormFile("image")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to read file"))
		return
	}
	defer image.Close()

	imagePath := "uploads/partners/images/" + imageHeader.Filename

	outImage, err := os.Create(imagePath)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to save file"))
		return
	}
	defer outImage.Close()

	_, err = io.Copy(outImage, image)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image: %v", err))
		return
	}

	partner := types_common_admin.PartnersPayload{
		Image: imagePath,
	}

	err = h.store.CreatePartner(&partner)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

// @Summary list partners
// @Description list partners
// @Tags common-admin
// @Accept json
// @Produce json
// @Router /admin/common/partner/list [get]
// @Security BearerAuth
func (h *Handler) handleListPartner(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.ListPartner()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, list)
}

// @Summart update partner
// @Description update partner
// @Tags common-admin
// @Accept multipart/json
// @Produce json
// @Param partnerId path string true "partner id"
// @Param image formData file true "image"
// @Router /admin/common/partner/{partnerId}/update [put]
// @Security BearerAuth
func (h *Handler) handleUpdatePartner(w http.ResponseWriter, r *http.Request) {
	var partnerId = r.PathValue("partnerId")

	partner, err := h.store.GetPartner(partnerId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = r.ParseMultipartForm(50)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	image, imageHeader, err := r.FormFile("image")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to read image"))
		return
	}
	defer image.Close()

	imagePath := "uploads/partners/images/" + imageHeader.Filename

	outImage, err := os.Create(imagePath)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to save file"))
		return
	}
	defer outImage.Close()

	_, err = io.Copy(outImage, image)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image"))
		return
	}

	changed_partner := types_common_admin.PartnersPayload{
		Image: imagePath,
	}

	err = h.store.UpdatePartner(partner.Id, &changed_partner)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "successfully updated"})
}

// @Summary delete partner
// @Description delete partner
// @Tags common-admin
// @Accept json
// @Produce json
// @Param partnerId path string true "partner id"
// @Router /admin/common/partner/{partnerId}/delete [delete]
// @Security BearerAuth
func (h *Handler) handleDeletePartner(w http.ResponseWriter, r *http.Request) {
	var partnerId = r.PathValue("partnerId")

	partner, err := h.store.GetPartner(partnerId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if partner == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("partner not found"))
		return
	}

	err = h.store.DeletePartner(partner.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": "successfully deleted"})
}

// @Summart get partner
// @Description get partner
// @Tags common-admin
// @Accept json
// @Produce json
// @Param partnerId path string true "partnerId"
// @Router /admin/common/partner/{partnerId} [get]
// @Security BearerAuth
func (h *Handler) handleGetPartner(w http.ResponseWriter, r *http.Request) {
	var partnerId = r.PathValue("partnerId")
	partner, err := h.store.GetPartner(partnerId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, partner)
}

// @Summart create banner
// @Description create banner
// @Tags common-admin
// @Accept multipart/data
// @Produce json
// @Param imageUz formData file true "image uz"
// @Param imageRu formData file true "image ru"
// @Param imageEn formData file true "image en"
// @Router /admin/common/banner/create [post]
// @Security BearerAuth
func (h *Handler) handleCreateBanner(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(50)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	imageUz, imageUzHeader, err := r.FormFile("imageUz")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unabel to read image_uz: %v", err))
		return
	}
	defer imageUz.Close()

	imageRu, imageRuHeader, err := r.FormFile("imageRu")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unabel to read image_ru: %v", err))
		return
	}
	defer imageRu.Close()

	imageEn, imageEnHeader, err := r.FormFile("imageEn")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unabel to read image_en: %v", err))
		return
	}
	defer imageEn.Close()

	imageUzPath := "uploads/banners/images/" + imageUzHeader.Filename
	imageRuPath := "uploads/banners/images/" + imageRuHeader.Filename
	imageEnPath := "uploads/banners/images/" + imageEnHeader.Filename

	outImageUz, err := os.Create(imageUzPath)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image_uz: %v", err))
		return
	}
	defer outImageUz.Close()

	outImageRu, err := os.Create(imageRuPath)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image_ru: %v", err))
		return
	}
	defer outImageRu.Close()

	outImageEn, err := os.Create(imageEnPath)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image_en: %v", err))
		return
	}
	defer outImageEn.Close()

	_, err = io.Copy(outImageUz, imageUz)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image_uz: %v", err))
		return
	}

	_, err = io.Copy(outImageRu, imageRu)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image_ru: %v", err))
		return
	}

	_, err = io.Copy(outImageEn, imageEn)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image_en: %v", err))
		return
	}

	banner := types_common_admin.BannerPayload{
		ImageUz: imageUzPath,
		ImageRu: imageRuPath,
		ImageEn: imageEnPath,
	}

	new_banner, err := h.store.CreateBanner(&banner)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, new_banner)
}

// @Summary list banner
// @Description list banner
// @Tags common-admin
// @Accept json
// @Produce json
// @Router /admin/common/banner/list [get]
// @Security BearerAuth
func (h *Handler) handleListBanner(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.ListBanner()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, list)
}

// @Summary update banner
// @Description update banner
// @Tags common-admin
// @Accept multipart/data
// @Produce json
// @Param bannerId path string true "banner id"
// @Param imageUz formData file true "image uz"
// @Param imageRu formData file true "image ru"
// @Param imageEn formData file true "image en"
// @Router /admin/common/banner/{bannerId}/update [put]
// @Security BearerAuth
func (h *Handler) handleUpdateBanner(w http.ResponseWriter, r *http.Request) {
	var bannerId = r.PathValue("bannerId")
	banner, err := h.store.GetBanner(bannerId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if banner == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("banner not found"))
		return
	}

	err = r.ParseMultipartForm(50)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	imageUz, imageUzHeader, err := r.FormFile("imageUz")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unabel to read image_uz: %v", err))
		return
	}
	defer imageUz.Close()

	imageRu, imageRuHeader, err := r.FormFile("imageRu")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unabel to read image_ru: %v", err))
		return
	}
	defer imageRu.Close()

	imageEn, imageEnHeader, err := r.FormFile("imageEn")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unabel to read image_en: %v", err))
		return
	}
	defer imageEn.Close()

	imageUzPath := "uploads/banners/images/" + imageUzHeader.Filename
	imageRuPath := "uploads/banners/images/" + imageRuHeader.Filename
	imageEnPath := "uploads/banners/images/" + imageEnHeader.Filename

	outImageUz, err := os.Create(imageUzPath)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image_uz: %v", err))
		return
	}
	defer outImageUz.Close()

	outImageRu, err := os.Create(imageRuPath)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image_ru: %v", err))
		return
	}
	defer outImageRu.Close()

	outImageEn, err := os.Create(imageEnPath)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image_en: %v", err))
		return
	}
	defer outImageEn.Close()

	_, err = io.Copy(outImageUz, imageUz)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image_uz: %v", err))
		return
	}

	_, err = io.Copy(outImageRu, imageRu)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image_ru: %v", err))
		return
	}

	_, err = io.Copy(outImageEn, imageEn)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image_en: %v", err))
		return
	}

	changed_banner := types_common_admin.BannerPayload{
		ImageUz: imageUzPath,
		ImageRu: imageRuPath,
		ImageEn: imageEnPath,
	}
	updated_banner, err := h.store.UpdateBanner(banner.Id, &changed_banner)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, updated_banner)
}

// @Summary delete banner
// @Description delete banner
// @Tags common-admin
// @Accept json
// @Produce json
// @Param bannerId path string true "banner id"
// @Router /admin/common/banner/{bannerId}/delete [delete]
// @Security BearerAuth
func (h *Handler) handleDeleteBanner(w http.ResponseWriter, r *http.Request) {
	var bannerId = r.PathValue("bannerId")
	banner, err := h.store.GetBanner(bannerId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if banner == nil {
		utils.WriteJson(w, http.StatusNotFound, map[string]string{"message": "banner not found"})
		return
	}

	err = h.store.DeleteBanner(banner.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": "succesfully deleted"})
}

// @Summary get banner
// @Description get banner
// @Tags common-admin
// @Accept json
// @Produce json
// @Param bannerId path string true "banner id"
// @Router /admin/common/banner/{bannerId} [get]
// @Security BearerAuth
func (h *Handler) handleGetBanner(w http.ResponseWriter, r *http.Request) {
	var bannerId = r.PathValue("bannerId")
	banner, err := h.store.GetBanner(bannerId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if banner == nil {
		utils.WriteJson(w, http.StatusNotFound, map[string]string{"message": "not found"})
		return
	}

	utils.WriteJson(w, http.StatusOK, banner)
}

// @Summart create news
// @Description create news
// @Tags common-admin
// @Accept multipart/data
// @Produce json
// @Param titleUz formData string true "title uz"
// @Param titleRu formData string true "title ru"
// @Param titleEn formData string true "title en"
// @Param descriptionUz formData string true "description uz"
// @Param descriptionRu formData string true "description ru"
// @Param descriptionEn formData string true "description en"
// @Param image formData file true "image"
// @Param link formData string false "link"
// @Router /admin/common/news/create [post]
// @Security BearerAuth
func (h *Handler) handleCreateNews(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(50)
	if err != nil {
		utils.WriteError(w, http.StatusBadGateway, err)
		return
	}

	titleUz := r.FormValue("titleUz")
	titleRu := r.FormValue("titleRu")
	titleEn := r.FormValue("titleEn")
	descriptionUz := r.FormValue("descriptionUz")
	descriptionRu := r.FormValue("descriptionRu")
	descriptionEn := r.FormValue("descriptionEn")
	link := r.FormValue("link")
	image, imageHeader, err := r.FormFile("image")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to read image: %v", err))
		return
	}
	defer image.Close()
	imagePath := "uploads/news/images/" + imageHeader.Filename

	outImage, err := os.Create(imagePath)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image: %v", err))
		return
	}
	defer outImage.Close()

	_, err = io.Copy(outImage, image)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image: %v", err))
		return
	}

	payload := types_common_admin.NewsPayload{
		TitleUz:       titleUz,
		TitleRu:       titleRu,
		TitleEn:       titleEn,
		DescriptionUz: descriptionUz,
		DescriptionRu: descriptionRu,
		DescriptionEn: descriptionEn,
		Image:         imagePath,
		Link:          link,
	}

	news, err := h.store.CreateNews(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, news)
}

// @Summary list news
// @Description list news
// @Tags common-admin
// @Accept json
// @Produce json
// @Param limit query string false "limit"
// @Param page query string false "page"
// @Router /admin/common/news/list [get]
// @Security BearerAuth
func (h *Handler) handleListNews(w http.ResponseWriter, r *http.Request) {
	var limit int = 10
	var offset int = 0

	query := r.URL.Query()
	if query.Get("limit") != "" {
		parsedLimit, err := strconv.Atoi(query.Get("limit"))
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	if query.Get("page") != "" {
		parsedPage, err := strconv.Atoi(query.Get("page"))
		if err == nil && parsedPage > 0 {
			offset = (parsedPage - 1) * limit
		}
	}

	list, err := h.store.ListNews(limit, offset)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	utils.WriteJson(w, http.StatusOK, map[string]interface{}{
		"page":  offset/limit + 1,
		"limit": limit,
		"news":  list,
	})
}

// @Summart update news
// @Description update news
// @Tags common-admin
// @Accept multipart/data
// @Produce json
// @Param newsId path string true "news Id"
// @Param titleUz formData string true "title uz"
// @Param titleRu formData string true "title ru"
// @Param titleEn formData string true "title en"
// @Param descriptionUz formData string true "description uz"
// @Param descriptionRu formData string true "description ru"
// @Param descriptionEn formData string true "description en"
// @Param image formData file true "image"
// @Param link formData string false "link"
// @Router /admin/common/news/{newsId}/update [put]
// @Security BearerAuth
func (h *Handler) handleUpdateNews(w http.ResponseWriter, r *http.Request) {
	var newsId = r.PathValue("newsId")
	news, err := h.store.GetNews(newsId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if news == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	err = r.ParseMultipartForm(50)
	if err != nil {
		utils.WriteError(w, http.StatusBadGateway, err)
		return
	}

	titleUz := r.FormValue("titleUz")
	titleRu := r.FormValue("titleRu")
	titleEn := r.FormValue("titleEn")
	descriptionUz := r.FormValue("descriptionUz")
	descriptionRu := r.FormValue("descriptionRu")
	descriptionEn := r.FormValue("descriptionEn")
	link := r.FormValue("link")
	image, imageHeader, err := r.FormFile("image")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to read image: %v", err))
		return
	}
	defer image.Close()
	imagePath := "uploads/news/images/" + imageHeader.Filename

	outImage, err := os.Create(imagePath)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image: %v", err))
		return
	}
	defer outImage.Close()

	_, err = io.Copy(outImage, image)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image: %v", err))
		return
	}

	payload := types_common_admin.NewsPayload{
		TitleUz:       titleUz,
		TitleRu:       titleRu,
		TitleEn:       titleEn,
		DescriptionUz: descriptionUz,
		DescriptionRu: descriptionRu,
		DescriptionEn: descriptionEn,
		Image:         imagePath,
		Link:          link,
	}

	new, err := h.store.UpdateNews(news.Id, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, new)
}

// @Summary delete news
// @Description delete news
// @Tags common-admin
// @Accept json
// @Produce json
// @Param newsId path string true "news Id"
// @Router /admin/common/news/{newsId}/delete [delete]
// @Security BearerAuth
func (h *Handler) handleDeleteNews(w http.ResponseWriter, r *http.Request) {
	var newsId = r.PathValue("newsId")
	log.Println(newsId)
	news, err := h.store.GetNews(newsId)
	log.Println(news)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if news == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	if err := h.store.DeleteNews(news.Id); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": "deleted"})
}

// @Summary get news
// @Description get news
// @Tags common-admin
// @Accept json
// @Produce json
// @Param newsId path string true "news Id"
// @Router /admin/common/news/{newsId} [get]
// @Security BearerAuth
func (h *Handler) handleGetNews(w http.ResponseWriter, r *http.Request) {
	var newsId = r.PathValue("newsId")
	news, err := h.store.GetNews(newsId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, news)
}

// @Summary create certificate
// @Description create certificate
// @Tags common-admin
// @Accept json
// @Produce json
// @Param payload body types_common_admin.CertificatePayload true "payload"
// @Router /admin/common/certificate/create [post]
// @Security BearerAuth
func (h *Handler) handleCreateCertificate(w http.ResponseWriter, r *http.Request) {
	var payload types_common_admin.CertificatePayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, errors)
		return
	}

	certificate := types_common_admin.CertificatePayload{
		NameUz: payload.NameUz,
		NameRu: payload.NameRu,
		NameEn: payload.NameEn,
		TextUz: payload.TextUz,
		TextRu: payload.TextRu,
		TextEn: payload.TextEn,
	}

	result, err := h.store.CreateCertificate(&certificate)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, result)
}

// @Summary list certificate
// @Description list certificate
// @Tags common-admin
// @Accept json
// @Produce json
// @Router /admin/common/certificate/list [get]
// @Security BearerAuth
func (h *Handler) handleListCertificate(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.ListCertificate()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, list)
}

// @Summary delete certificate
// @Summary delete certificate
// @Tags common-admin
// @Accept json
// @Produce json
// @Param certificateId path string true "certificate id"
// @Router /admin/common/certificate/{certificateId}/delete [delete]
// @Security BearerAuth
func (h *Handler) handleDeleteCertificate(w http.ResponseWriter, r *http.Request) {
	var certificateId = r.PathValue("certificateId")
	certificate, err := h.store.GetCertificate(certificateId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if certificate == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	err = h.store.DeleteCertificate(certificate.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": "deleted"})
}

// @Summary update certificate
// @Description update certificate
// @Tags common-admin
// @Accept json
// @Produce json
// @Param certificateId path string true "certificate id"
// @Param payload body types_common_admin.CertificatePayload true "payload"
// @Router /admin/common/certificate/{certificateId}/update [put]
// @Security BearerAuth
func (h *Handler) handleUpdateCertificate(w http.ResponseWriter, r *http.Request) {
	var certificateId = r.PathValue("certificateId")
	certificate, err := h.store.GetCertificate(certificateId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if certificate == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	var payload types_common_admin.CertificatePayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, errors)
		return
	}

	updatedCertificate, err := h.store.UpdateCertificate(certificate.Id, &types_common_admin.CertificatePayload{
		NameUz: payload.NameUz,
		NameRu: payload.NameRu,
		NameEn: payload.NameEn,
		TextUz: payload.TextUz,
		TextRu: payload.TextRu,
		TextEn: payload.TextEn,
	})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, updatedCertificate)
}

// @Summary get certificate
// @Description get certificate
// @Tags common-admin
// @Accept json
// @Produce json
// @Param certificateId path string true "certificate id"
// @Router /admin/common/certificate/{certificateId} [get]
// @Security BearerAuth
func (h *Handler) handleGetCertificate(w http.ResponseWriter, r *http.Request) {
	var certificateId = r.PathValue("certificateId")
	certificate, err := h.store.GetCertificate(certificateId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, certificate)
}
