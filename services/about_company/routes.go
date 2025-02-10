package about_company

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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
	// capasity
	r.Post("/admin/company/capasity/create", auth.AuthWithJWT(h.handleCreateCapasity, h.userStore))
	r.Get("/admin/company/capasity/list", auth.AuthWithJWT(h.handleListCapasity, h.userStore))
	r.Delete("/admin/company/capasity/{capasityId}/delete", auth.AuthWithJWT(h.handleDeleteCapasity, h.userStore))
	r.Put("/admin/company/capasity/{capasityId}/update", auth.AuthWithJWT(h.handleUpdateCapasity, h.userStore))
	r.Get("/admin/company/capasity/{capasityId}", auth.AuthWithJWT(h.handleGetCapasity, h.userStore))
	// about oil
	r.Post("/admin/company/about_oil/create", auth.AuthWithJWT(h.handleCreateAboutOil, h.userStore))
	r.Get("/admin/company/about_oil/list", auth.AuthWithJWT(h.handleListAboutOil, h.userStore))
	r.Put("/admin/company/about_oil/{oilId}/update", auth.AuthWithJWT(h.handleUpdateAboutOil, h.userStore))
	r.Delete("/admin/company/about_oil/{oilId}/delete", auth.AuthWithJWT(h.handleDeleteAboutOil, h.userStore))
	r.Get("/admin/company/about_oil/{oilId}", auth.AuthWithJWT(h.handleGetAboutOil, h.userStore))
	// why us
	r.Post("/admin/company/why_us/create", auth.AuthWithJWT(h.handleCreateWhyUs, h.userStore))
	r.Get("/admin/company/why_us/list", auth.AuthWithJWT(h.handleListWhyUs, h.userStore))
	r.Put("/admin/company/why_us/{id}/update", auth.AuthWithJWT(h.handleUpdateWhyUs, h.userStore))
	r.Delete("/admin/company/why_us/{id}/delete", auth.AuthWithJWT(h.handleDeleteWhyUs, h.userStore))
	r.Get("/admin/company/why_us/{id}", auth.AuthWithJWT(h.handleGetWhyUs, h.userStore))
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

// @Summary h=create about oil
// @DEscription create about oil
// @Tags company-admin
// @Accept json
// @Produce json
// @Param payload body types_about_company.AboutOilPayload true "payload"
// @Router /admin/company/about_oil/create [post]
// @Security BearerAuth
func (h *Handler) handleCreateAboutOil(w http.ResponseWriter, r *http.Request) {
	var payload types_about_company.AboutOilPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	about_oil, err := h.store.CreateAboutOil(&types_about_company.AboutOilPayload{
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

	utils.WriteJson(w, http.StatusCreated, about_oil)
}

// @Summary list about oil
// @Description about oil
// @Tags company-admin
// @Accept json
// @Produce json
// @Router /admin/company/about_oil/list [get]
// @Security BearerAuth
func (h *Handler) handleListAboutOil(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.ListAboutOil()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, list)
}

// @Summary update about oil
// @Description update about oil
// @Tags company-admin
// @Accept json
// @Produce json
// @Param oilId path string true "id"
// @Param payload body types_about_company.AboutOilPayload true "payload"
// @Router /admin/company/about_oil/{oilId}/update [put]
// @Security BearerAuth
func (h *Handler) handleUpdateAboutOil(w http.ResponseWriter, r *http.Request) {
	var oilId = r.PathValue("oilId")
	oil, err := h.store.GetAboutOil(oilId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if oil == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}
	var payload types_about_company.AboutOilPayload

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, errors)
		return
	}

	updated_oil := types_about_company.AboutOilPayload{
		NameUz: payload.NameUz,
		NameRu: payload.NameRu,
		NameEn: payload.NameEn,
		TextUz: payload.TextUz,
		TextRu: payload.TextRu,
		TextEn: payload.TextEn,
	}

	about_oil, err := h.store.UpdateAboutOil(oil.Id, &updated_oil)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, about_oil)
}

// @Summary delete about oil
// @Description delete about oil
// @Tags company-admin
// @Accept json
// @Produce json
// @Param oilId path string true "id"
// @Router /admin/company/about_oil/{oilId}/delete [delete]
// @Security BearerAuth
func (h *Handler) handleDeleteAboutOil(w http.ResponseWriter, r *http.Request) {
	var oilId = r.PathValue("oilId")
	oil, err := h.store.GetAboutOil(oilId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if oil == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	err = h.store.DeleteAboutOil(oil.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": "deleted"})
}

// @Summary get about oil
// @Description get about oil
// @Tags company-admin
// @Accept json
// @Produce json
// @Param oilId path string true "id"
// @Router /admin/company/about_oil/{oilId} [get]
// @Security BearerAuth
func (h *Handler) handleGetAboutOil(w http.ResponseWriter, r *http.Request) {
	var oilId = r.PathValue("oilId")
	oil, err := h.store.GetAboutOil(oilId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, oil)
}

// @Summary create why us
// @Description create why us
// @Tags company-admin
// @Accept multipart/data
// @Produce json
// @Param image formData file false "image"
// @Param titleUz formData string true "titleUz"
// @Param titleRu formData string true "titleRu"
// @Param titleEn formData string true "titleEn"
// @Param descriptionUz formData string true "descriptionUz"
// @Param descriptionRu formData string true "descriptionRu"
// @Param descriptionEn formData string ture "descriptionEn"
// @Router /admin/company/why_us/create [post]
// @Security BearerAuth
func (h *Handler) handleCreateWhyUs(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(50)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	titleUz := r.FormValue("titleUz")
	titleRu := r.FormValue("titleRu")
	titleEn := r.FormValue("titleEn")
	descriptionUz := r.FormValue("descriptionUz")
	descriptionRu := r.FormValue("descriptionRu")
	descriptionEn := r.FormValue("descriptionEn")
	image, imageHeader, err := r.FormFile("image")

	if err != nil {
		if err == http.ErrMissingFile {
			image = nil
			imageHeader = nil
		} else {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to read image: %v", err))
			return
		}
	}
	var imagePath string
	var outImage *os.File
	if image != nil {
		defer image.Close()
		if imageHeader != nil {
			imagePath = "uploads/why_us/images/" + imageHeader.Filename
		} else {
			imagePath = ""
		}
		if imagePath != "" {
			outImage, err = os.Create(imagePath)
			if err != nil {
				utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to save image: %v", err))
				return
			}
			defer outImage.Close()
			_, err = io.Copy(outImage, image)
			if err != nil {
				utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image: %v", err))
				return
			}
		}
	} else {
		imagePath = ""
	}

	whyUs := types_about_company.WhyUsPayload{
		TitleUz:       titleUz,
		TitleRu:       titleRu,
		TitleEn:       titleEn,
		DescriptionUz: descriptionUz,
		DescriptionRu: descriptionRu,
		DescriptionEn: descriptionEn,
		Image:         imagePath,
	}

	result, err := h.store.CreateWhyUs(&whyUs)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, result)
}

// @Summary list why_us
// @Description list why_us
// @Tags company-admin
// @Accept json
// @Produce json
// @Router /admin/company/why_us/list [get]
// @Security BearerAuth
func (h *Handler) handleListWhyUs(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.ListWhyUs()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	utils.WriteJson(w, http.StatusOK, list)
}

// @Summary get why us
// @Summary get why us
// @Tags company-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Router /admin/company/why_us/{id} [get]
// @Security BearerAuth
func (h *Handler) handleGetWhyUs(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	whyUs, err := h.store.GetWhyUs(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, whyUs)
}

// @Summary delete why us
// @Description delete why us
// @Tags company-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Router /admin/company/why_us/{id}/delete [delete]
// @Security BearerAuth
func (h *Handler) handleDeleteWhyUs(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	whyUs, err := h.store.GetWhyUs(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if whyUs == nil {
		utils.WriteJson(w, http.StatusNotFound, map[string]string{"message": "not found"})
		return
	}

	err = h.store.DeleteWhyUs(whyUs.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": "deleted"})
}

// @Summary update why us
// @Description update why us
// @Tags company-admin
// @Accpet multipart/data
// @Produce json
// @Param id path string true "id"
// @Param image formData file false "image"
// @Param titleUz formData string false "titleUz"
// @Param titleRu formData string false "titleRu"
// @Param titleEn formData string false "titleEn"
// @Param descriptionUz formData string false "descriptionUz"
// @Param descriptionRu formData string false "descriptionRu"
// @Param descriptionEn formData string false "descriptionEn"
// @Router /admin/company/why_us/{id}/update [put]
// @Security BearerAuth
func (h *Handler) handleUpdateWhyUs(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	whyUs, err := h.store.GetWhyUs(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	titleUz := r.FormValue("titleUz")
	titleRu := r.FormValue("titleRu")
	titleEn := r.FormValue("titleEn")
	descriptionUz := r.FormValue("descriptionUz")
	descriptionRu := r.FormValue("descriptionRu")
	descriptionEn := r.FormValue("descriptionEn")
	image, imageHeader, err := r.FormFile("image")

	if err != nil {
		if err == http.ErrMissingFile {
			image = nil
			imageHeader = nil
		} else {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to read image: %v", err))
			return
		}
	}
	var imagePath string
	var outImage *os.File
	if image != nil {
		defer image.Close()
		if imageHeader != nil {
			imagePath = "uploads/why_us/images/" + imageHeader.Filename
		} else {
			imagePath = ""
		}
		if imagePath != "" {
			outImage, err = os.Create(imagePath)
			if err != nil {
				utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to save image: %v", err))
				return
			}
			defer outImage.Close()
			_, err = io.Copy(outImage, image)
			if err != nil {
				utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image: %v", err))
				return
			}
		}
	} else {
		imagePath = ""
	}

	changedWhyUs := types_about_company.WhyUsPayload{
		TitleUz:       titleUz,
		TitleRu:       titleRu,
		TitleEn:       titleEn,
		DescriptionUz: descriptionUz,
		DescriptionRu: descriptionRu,
		DescriptionEn: descriptionEn,
		Image:         imagePath,
	}
	log.Println("next")
	result, err := h.store.UpdateWhyUs(whyUs.Id, &changedWhyUs)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		log.Println("next")
		return
	}

	utils.WriteJson(w, http.StatusOK, result)
}
