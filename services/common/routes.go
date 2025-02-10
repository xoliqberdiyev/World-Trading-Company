package common

import (
	"fmt"
	"net/http"
	"strconv"

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
	r.Get("/category/list", h.handleListCategory)
	r.Get("/banner/list", h.handleListBanner)
	r.Get("/news/list", h.handleListNews)
	r.Get("/news/{newsId}", h.handleGetNews)
	r.Get("/about_oil/list", h.handleListAboutOil)
	r.Get("/certificate/list", h.handleListCertificate)
	r.Get("/why_us/list", h.handleListWhyUs)
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

// @Summary list category
// @Description list category
// @Tags common
// @Accept json
// @Produce json
// @Router /category/list [get]
func (h *Handler) handleListCategory(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.ListCategory()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, list)
}

// @Summart list banner
// @Decription list banner
// @Tags common
// @Accept json
// @Produce json
// @Router /banner/list [get]
func (h *Handler) handleListBanner(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.ListBanner()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, list)
}

// @Summary get news
// @Description get news
// @Tags common
// @Accept json
// @Produce json
// @Param newsId path string true "news Id"
// @Router /news/{newsId} [get]
func (h *Handler) handleGetNews(w http.ResponseWriter, r *http.Request) {
	var newsId = r.PathValue("newsId")
	news, err := h.store.GetNews(newsId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, news)
}

// @Summary list news
// @Description list news
// @Tags common
// @Accept json
// @Produce json
// @Param limit query string false "limit"
// @Param page query string false "page"
// @Router /news/list [get]
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

// @Summary list about oil
// @Description list about oil
// @Tags common
// @Accept json
// @Produce json
// @Router /about_oil/list [get]
func (h *Handler) handleListAboutOil(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.ListAboutOil()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, list)
}

// @Summary list certificate
// @Description list certificate
// @Tags common
// @Accept json
// @Produce json
// @Router /certificate/list [get]
func (h *Handler) handleListCertificate(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.ListCertificate()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, list)
}

// @Summary list why us
// @Description list why us
// @Tags common
// @Accept json
// @Produce json
// @Router /why_us/list [get]
func (h *Handler) handleListWhyUs(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.ListWhyUs()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, list)
}
