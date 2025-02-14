package product

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/XoliqberdiyevBehruz/wtc_backend/services/auth"
	types_product "github.com/XoliqberdiyevBehruz/wtc_backend/types/product"
	types_user_admin "github.com/XoliqberdiyevBehruz/wtc_backend/types/user_admin"
	"github.com/XoliqberdiyevBehruz/wtc_backend/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
)

type Handler struct {
	store     types_product.ProductStore
	userStore types_user_admin.UserStore
}

func NewHandler(store types_product.ProductStore, userStore types_user_admin.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegsiterRoutes(r chi.Router) {
	// category
	r.Get("/admin/product/category/list", auth.AuthWithJWT(h.handleListCategory, h.userStore))
	r.Post("/admin/product/category/create", auth.AuthWithJWT(h.handleCreateCategory, h.userStore))
	r.Put("/admin/product/crategory/{categoryId}/update", auth.AuthWithJWT(h.handleUpdateCategory, h.userStore))
	r.Delete("/admin/product/category/{categoryId}/delete", auth.AuthWithJWT(h.handleDeleteCategory, h.userStore))
	r.Get("/admin/product/category/{categoryId}", auth.AuthWithJWT(h.handleGetCategory, h.userStore))
	// product
	r.Post("/admin/product/product/create", auth.AuthWithJWT(h.handleCreateProduct, h.userStore))
	r.Get("/admin/product/product/list", auth.AuthWithJWT(h.handleListProduct, h.userStore))
	r.Get("/admin/product/product/{id}", auth.AuthWithJWT(h.handleGetProduct, h.userStore))
	r.Delete("/admin/product/product/{id}/delete", auth.AuthWithJWT(h.handleDeleteProduct, h.userStore))
	r.Put("/admin/product/product/{id}/update", auth.AuthWithJWT(h.handleUpdateProduct, h.userStore))
	// product media
	r.Post("/admin/product/product_media/create", auth.AuthWithJWT(h.handleCreateProductMedia, h.userStore))
	r.Get("/admin/product/product_media/list", auth.AuthWithJWT(h.handleListProductMedia, h.userStore))
	r.Get("/admin/product/product_media/{id}", auth.AuthWithJWT(h.handleGetProductMedia, h.userStore))
	r.Delete("/admin/product/product_media/{id}/delete", auth.AuthWithJWT(h.handleDeleteProductMedia, h.userStore))
	r.Put("/admin/product/product_media/{id}/update", auth.AuthWithJWT(h.handleUpdateProductMedia, h.userStore))
	// product spesification
	r.Post("/admin/product/spesification/create", auth.AuthWithJWT(h.handleCreateProductSpesification, h.userStore))
	r.Get("/admin/product/spesification/list", auth.AuthWithJWT(h.handleListProductSpesification, h.userStore))
	r.Get("/admin/product/spesification/{id}", auth.AuthWithJWT(h.handleGetProductSpesification, h.userStore))
	r.Delete("/admin/product/spesification/{id}/delete", auth.AuthWithJWT(h.handleDeleteProductSpesification, h.userStore))
	r.Put("/admin/product/spesification/{id}/update", auth.AuthWithJWT(h.handleUpdateProductSpesification, h.userStore))
	// product features
	r.Post("/admin/product/feature/create", auth.AuthWithJWT(h.handleCreateProductFeature, h.userStore))
	r.Get("/admin/product/feature/list", auth.AuthWithJWT(h.handleListProductFeature, h.userStore))
	r.Get("/admin/product/feature/{id}", auth.AuthWithJWT(h.handleGetProductFeature, h.userStore))
	r.Delete("/admin/product/feature/{id}/delete", auth.AuthWithJWT(h.handleDeleteProductFeature, h.userStore))
	r.Put("/admin/product/feature/{id}/update", auth.AuthWithJWT(h.handleUpdateProductFeature, h.userStore))
	// product adventage
	r.Post("/admin/product/adventage/create", auth.AuthWithJWT(h.handleCreateProductAdventage, h.userStore))
	r.Get("/admin/product/adventage/list", auth.AuthWithJWT(h.handleListProductAdventage, h.userStore))
	r.Get("/admin/product/adventage/{id}", auth.AuthWithJWT(h.handleGetProductAdventage, h.userStore))
	r.Delete("/admin/product/adventage/{id}/delete", auth.AuthWithJWT(h.handleDeleteProductAdventage, h.userStore))
	r.Put("/admin/product/adventage/{id}/update", auth.AuthWithJWT(h.handleUpdateProductAdventage, h.userStore))
	// product chemistry
	r.Post("/admin/product/chemistry/create", auth.AuthWithJWT(h.handleCreateChemistry, h.userStore))
	r.Get("/admin/product/chemistry/list", auth.AuthWithJWT(h.handleListChemistry, h.userStore))
	r.Get("/admin/product/chemistry/{id}", auth.AuthWithJWT(h.handleGetChemistry, h.userStore))
	r.Delete("/admin/product/chemistry/{id}/delete", auth.AuthWithJWT(h.handleDeleteChemistry, h.userStore))
	r.Put("/admin/product/chemistry/{id}/update", auth.AuthWithJWT(h.handleUpdateChemistry, h.userStore))
	// product impact
	r.Post("/admin/product/impact/create", auth.AuthWithJWT(h.handleCreateImpact, h.userStore))
	r.Get("/admin/product/impact/list", auth.AuthWithJWT(h.handleListImpact, h.userStore))
	r.Get("/admin/product/impact/{id}", auth.AuthWithJWT(h.handleGetImpact, h.userStore))
	r.Delete("/admin/product/impact/{id}/delete", auth.AuthWithJWT(h.handleDeleteImpact, h.userStore))
	r.Put("/admin/product/impact/{id}/update", auth.AuthWithJWT(h.handleUpdateImpact, h.userStore))
	// product files
	r.Post("/admin/product/file/create", auth.AuthWithJWT(h.handleCreateProductFile, h.userStore))
	r.Get("/admin/product/file/list", auth.AuthWithJWT(h.handleListProductFile, h.userStore))
	r.Get("/admin/product/file/{id}", auth.AuthWithJWT(h.handleGetProductFile, h.userStore))
	r.Delete("/admin/product/file/{id}/delete", auth.AuthWithJWT(h.handleDeleteProductFile, h.userStore))
	r.Put("/admin/product/file/{id}/update", auth.AuthWithJWT(h.handleUpdateProductFile, h.userStore))
}

// @Summary create category
// @Description create category
// @Tags product-admin
// @Accept multipart/data
// @Produce json
// @Param nameUz formData string true "name uz"
// @Param nameRu formData string true "name ru"
// @Param nameEn formData string true "name en"
// @Param image formData file true "image"
// @Param icon formData file true "icon"
// @Router /admin/product/category/create [post]
// @Security BearerAuth
func (h *Handler) handleCreateCategory(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(50)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	nameUz := r.FormValue("nameUz")
	nameRu := r.FormValue("nameRu")
	nameEn := r.FormValue("nameEn")

	image, imageHeader, err := r.FormFile("image")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to read image: %v", err))
		return
	}

	icon, iconHeader, err := r.FormFile("icon")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to read icon: %v", err))
		return
	}

	imagePath := "uploads/categories/images/" + imageHeader.Filename
	iconPath := "uploads/categories/icons/" + iconHeader.Filename

	outImage, err := os.Create(imagePath)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to copy image: %v", err))
		return
	}
	outIcon, err := os.Create(iconPath)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to copy icon: %v", err))
		return
	}

	_, err = io.Copy(outImage, image)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image: %v", err))
		return
	}

	_, err = io.Copy(outIcon, icon)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write icon: %v", err))
		return
	}

	category := types_product.CategoryPayload{
		NameUz: nameUz,
		NameRu: nameRu,
		NameEn: nameEn,
		Image:  imagePath,
		Icon:   iconPath,
	}

	err = h.store.CreateCategory(&category)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to create category: %v", err))
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

// @Summary list category
// @Description list category
// @Tags product-admin
// @Accept json
// @Produce json
// @Router /admin/product/category/list [get]
// @Security BearerAuth
func (h *Handler) handleListCategory(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.ListCategory()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, list)
}

// @Summary update category
// @Desctiption update category
// @Tags product-admin
// @Accept multipart/data
// @Produce json
// @Param categoryId path string true "category id"
// @Param nameUz formData string false "name uz"
// @Param nameRu formData string false "name ru"
// @Param nameEn formData string false "name en"
// @Param image formData file false "image"
// @Param icon formData file false "icon"
// @Router /admin/product/crategory/{categoryId}/update [put]
// @Security BearerAuth
func (h *Handler) handleUpdateCategory(w http.ResponseWriter, r *http.Request) {
	var categoryId = r.PathValue("categoryId")
	category, err := h.store.GetCategory(categoryId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if category == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("category not found"))
		return
	}

	err = r.ParseMultipartForm(50)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	nameUz := r.FormValue("nameUz")
	nameRu := r.FormValue("nameRu")
	nameEn := r.FormValue("nameEn")

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
	if image != nil {
		defer image.Close()
	}
	icon, iconHeader, err := r.FormFile("icon")
	if err != nil {
		if err == http.ErrMissingFile {
			icon = nil
			iconHeader = nil
		} else {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to read icon: %v", err))
			return
		}
	}
	if icon != nil {
		defer icon.Close()
	}
	var imagePath string
	var iconPath string
	if imageHeader != nil {
		imagePath = "uploads/categories/images/" + imageHeader.Filename
	}
	if iconHeader != nil {
		iconPath = "uploads/categories/icons/" + iconHeader.Filename
	}
	if imagePath != "" {
		outImage, err := os.Create(imagePath)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to copy image: %v", err))
			return
		}
		defer outImage.Close()

		_, err = io.Copy(outImage, image)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write image: %v", err))
			return
		}
	}
	if iconPath != "" {
		outIcon, err := os.Create(iconPath)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to copy icon: %v", err))
			return
		}
		defer outIcon.Close()
		_, err = io.Copy(outIcon, icon)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("unable to write icon: %v", err))
			return
		}
	}

	changed_category := types_product.CategoryPayload{
		NameUz: nameUz,
		NameRu: nameRu,
		NameEn: nameEn,
		Image:  imagePath,
		Icon:   iconPath,
	}

	err = h.store.UpdateCategory(category.Id, &changed_category)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "successfully updated"})
}

// @Summary delete category
// @Description delete category
// @Tags product-admin
// @Accept json
// @Produce json
// @Param categoryId path string true "category id"
// @Router /admin/product/category/{categoryId}/delete [delete]
// @Security BearerAuth
func (h *Handler) handleDeleteCategory(w http.ResponseWriter, r *http.Request) {
	var categoryId = r.PathValue("categoryId")
	category, err := h.store.GetCategory(categoryId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.DeleteCategory(category.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": "successfully deleted"})
}

// @Summary get category
// @Description get category
// @Tags product-admin
// @Acceept json
// @Produce json
// @Router /admin/product/category/{categoryId} [get]
// @Param categoryId path string true "category id"
// @Security BearerAuth
func (h *Handler) handleGetCategory(w http.ResponseWriter, r *http.Request) {
	var categoryId = r.PathValue("categoryId")
	category, err := h.store.GetCategory(categoryId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, category)
}

// @Summary create product
// @Description create product
// @Tags product-admin
// @Accelt multipart/data
// @Produce json
// @Param nameUz formData string true "name uz"
// @Param nameRu formData string true "name ru"
// @Param nameEn formData string true "name en"
// @Param descriptionUz formData string true "description uz"
// @Param descriptionRu formData string true "description ru"
// @Param descriptionEn formData string true "description en"
// @Param textUz formData string true "text uz"
// @Param textRu formData string true "text ru"
// @Param textEn formData string true "text en"
// @Param categoryId formData string true "category id"
// @Param image formData file true "image"
// @Param banner formData file true "banner"
// @Router /admin/product/product/create [post]
// @Security BearerAuth
func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(50)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// name
	nameUz := r.FormValue("nameUz")
	nameRu := r.FormValue("nameRu")
	nameEn := r.FormValue("nameEn")
	// description
	descriptionUz := r.FormValue("descriptionUz")
	descriptionRu := r.FormValue("descriptionRu")
	descriptionEn := r.FormValue("descriptionEn")
	// text
	textUz := r.FormValue("textUz")
	textRu := r.FormValue("textRu")
	textEn := r.FormValue("textEn")
	// category
	categoryId := r.FormValue("categoryId")
	category, err := h.store.GetCategory(categoryId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if category == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("category not found"))
		return
	}

	image, imageHeader, err := r.FormFile("image")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	defer image.Close()
	imagePath := "uploads/product/images/" + imageHeader.Filename
	outImage, err := os.Create(imagePath)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	defer outImage.Close()

	_, err = io.Copy(outImage, image)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	banner, bannerHeader, err := r.FormFile("banner")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	defer banner.Close()
	bannerPath := "uploads/product/banners/" + bannerHeader.Filename

	outBanner, err := os.Create(bannerPath)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	defer outBanner.Close()

	_, err = io.Copy(outBanner, banner)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	payload := types_product.ProductPayload{
		NameUz:        nameUz,
		NameRu:        nameRu,
		NameEn:        nameEn,
		DescriptionUz: descriptionUz,
		DescriptionRu: descriptionRu,
		DescriptionEn: descriptionEn,
		TextUz:        textUz,
		TextRu:        textRu,
		TextEn:        textEn,
		CategoryId:    category.Id,
		Image:         imagePath,
		Banner:        bannerPath,
	}

	result, err := h.store.CreateProduct(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, result)
}

// @Summary list product
// @Description list product
// @Tags product-admin
// @Accept json
// @Produce json
// @Param limit query string false "limit"
// @Param page query string false "page"
// @Router /admin/product/product/list [get]
// @Security BearerAuth
func (h *Handler) handleListProduct(w http.ResponseWriter, r *http.Request) {
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

	list, count, err := h.store.ListProduct(offset, limit)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]interface{}{
		"page":   offset/limit + 1,
		"limit":  limit,
		"count":  count,
		"result": list,
	})
}

// @Summary get product
// @Description get product
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Router /admin/product/product/{id} [get]
// @Security BearerAuth
func (h *Handler) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	product, err := h.store.GetProduct(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if product == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	utils.WriteJson(w, http.StatusOK, product)
}

// @Summary delete product
// @Description delete product
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Router /admin/product/product/{id}/delete [delete]
// @Security BearerAuth
func (h *Handler) handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	product, err := h.store.GetProduct(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if product == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	err = h.store.DeleteProduct(product.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": "deleted"})
}

// @Summary update product
// @Description update product
// @Tags product-admin
// @Accept multipart/data
// @Produce json
// @Param id path string true "id"
// @Param nameUz formData string false "name uz"
// @Param nameRu formData string false "name ru"
// @Param nameEn formData string false "name en"
// @Param descriptionUz formData string false "description uz"
// @Param descriptionRu formData string false "description ru"
// @Param descriptionEn formData string false "description en"
// @Param textUz formData string false "text uz"
// @Param textRu formData string false "text ru"
// @Param textEn formData string false "text en"
// @Param categoryId formData string false "category id"
// @Param image formData file false "image"
// @Param banner formData file false "banner"
// @Router /admin/product/product/{id}/update [put]
// @Security BearerAuth
func (h *Handler) handleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	product, err := h.store.GetProduct(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error to get product: %v", err))
		return
	}
	if product == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	err = r.ParseMultipartForm(50)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// name
	nameUz := r.FormValue("nameUz")
	nameRu := r.FormValue("nameRu")
	nameEn := r.FormValue("nameEn")
	// description
	descriptionUz := r.FormValue("descriptionUz")
	descriptionRu := r.FormValue("descriptionRu")
	descriptionEn := r.FormValue("descriptionEn")
	// text
	textUz := r.FormValue("textUz")
	textRu := r.FormValue("textRu")
	textEn := r.FormValue("textEn")
	// category
	categoryId := r.FormValue("categoryId")
	if categoryId != "" {
		category, err := h.store.GetCategory(categoryId)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
		if category == nil {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("category not found"))
			return
		}
	}

	image, imageHeader, err := r.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			image = nil
			imageHeader = nil
		} else {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
	}
	if image != nil {
		defer image.Close()
	}
	var imagePath string
	if imageHeader != nil {
		imagePath = "uploads/product/images/" + imageHeader.Filename
	} else {
		imagePath = ""
	}
	if imagePath != "" {
		outImage, err := os.Create(imagePath)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}

		defer outImage.Close()

		_, err = io.Copy(outImage, image)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
	}

	banner, bannerHeader, err := r.FormFile("banner")
	if err != nil {
		if err == http.ErrMissingFile {
			banner = nil
			bannerHeader = nil
		} else {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
	}
	if banner != nil {
		defer banner.Close()
	}
	var bannerPath string
	if bannerHeader != nil {
		bannerPath = "uploads/product/banners/" + bannerHeader.Filename
	} else {
		bannerPath = ""
	}
	if bannerPath != "" {
		outBanner, err := os.Create(bannerPath)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
		defer outBanner.Close()

		_, err = io.Copy(outBanner, banner)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
	}

	payload := types_product.ProductPayload{
		NameUz:        nameUz,
		NameRu:        nameRu,
		NameEn:        nameEn,
		DescriptionUz: descriptionUz,
		DescriptionRu: descriptionRu,
		DescriptionEn: descriptionEn,
		TextUz:        textUz,
		TextRu:        textRu,
		TextEn:        textEn,
		CategoryId:    categoryId,
		Image:         imagePath,
		Banner:        bannerPath,
	}

	result, err := h.store.UpdateProduct(product.Id, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, result)
}

// @Summary create product media
// @Description create product media
// @Tags product-admin
// @Accept multipart/data
// @Produce json
// @Param productId formData string true "product id"
// @Param image formData file true "image"
// @Router /admin/product/product_media/create [post]
// @Security BearerAuth
func (h *Handler) handleCreateProductMedia(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(50)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	productId := r.FormValue("productId")
	product, err := h.store.GetProduct(productId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if product == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	image, imageHeader, err := r.FormFile("image")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	defer image.Close()

	imagePath := "uploads/product_medias/images/" + imageHeader.Filename

	outImage, err := os.Create(imagePath)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	defer outImage.Close()

	_, err = io.Copy(outImage, image)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	payload := types_product.ProductMediaPayload{
		ProductId: product.Id,
		Image:     imagePath,
	}

	result, err := h.store.CreateProductMedia(payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, result)
}

// @Summary list product medias
// @Description list product medias
// @Tags product-admin
// @Accept json
// @Produce json
// @Param limit query string false "limit"
// @Param page query string false "page"
// @Router /admin/product/product_media/list [get]
// @Security BearerAuth
func (h *Handler) handleListProductMedia(w http.ResponseWriter, r *http.Request) {
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

	list, count, err := h.store.ListProductMedia(limit, offset)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]interface{}{
		"page":   offset/limit + 1,
		"limit":  limit,
		"count":  count,
		"result": list,
	})
}

// @Summary get product media
// @Descriotion get produce media
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Router /admin/product/product_media/{id} [get]
// @Security BearerAuth
func (h *Handler) handleGetProductMedia(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	product, err := h.store.GetProductMedia(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if product == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	utils.WriteJson(w, http.StatusOK, product)
}

// @Summary delete product_media
// @Description delete product media
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Router /admin/product/product_media/{id}/delete [delete]
// @Security BearerAuth
func (h *Handler) handleDeleteProductMedia(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	productMedia, err := h.store.GetProductMedia(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if productMedia == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	err = h.store.DeleteProductMedia(productMedia.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": "deleted"})
}

// @Summary update product media
// @Description update product media
// @Tags product-admin
// @Accept multipart/data
// @Produce json
// @Param id path string true "id"
// @Param productId formData string false "product id"
// @Param image formData file false "image"
// @Router /admin/product/product_media/{id}/update [put]
// @Security BearerAuth
func (h *Handler) handleUpdateProductMedia(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	productMedia, err := h.store.GetProductMedia(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if productMedia == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}
	err = r.ParseMultipartForm(50)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	productId := r.FormValue("productId")

	if productId != "" {
		product, err := h.store.GetProduct(productId)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
		if product == nil {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
			return
		}
	}
	image, imageHeader, err := r.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			image = nil
			imageHeader = nil
		} else {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
	}
	if image != nil {
		defer image.Close()
	}
	var imagePath string
	if imageHeader != nil {
		imagePath = "uploads/product_medias/images/" + imageHeader.Filename
	} else {
		imagePath = ""
	}
	if imagePath != "" {
		outImage, err := os.Create(imagePath)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
		defer outImage.Close()

		_, err = io.Copy(outImage, image)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
	}
	payload := types_product.ProductMediaPayload{
		Image:     imagePath,
		ProductId: productId,
	}

	result, err := h.store.UpdateProductMedia(productMedia.Id, payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, result)
}

// @Summary create product spesification
// @Description create product spesification
// @Tags product-admin
// @Accept json
// @Produce json
// @Param payload body types_product.ProductSpesificationPayload true "payload"
// @Router /admin/product/spesification/create [post]
// @Security BearerAuth
func (h *Handler) handleCreateProductSpesification(w http.ResponseWriter, r *http.Request) {
	var payload types_product.ProductSpesificationPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, errors)
		return
	}

	product, err := h.store.GetProduct(payload.ProductId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if product == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}

	err = h.store.CreateProductSpesification(types_product.ProductSpesificationPayload{
		NameUz:    payload.NameUz,
		NameRu:    payload.NameRu,
		NameEn:    payload.NameEn,
		Brands:    payload.Brands,
		ProductId: payload.ProductId,
	})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusCreated, map[string]string{"message": "created"})
}

// @Summary list product spesification
// @Description list product spesification
// @Tags product-admin
// @Accept json
// @Produce json
// @Router /admin/product/spesification/list [get]
// @Security BearerAuth
func (h *Handler) handleListProductSpesification(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.ListProductSpesification()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, list)
}

// @Summary get data
// @Description get data
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Router /admin/product/spesification/{id} [get]
// @Security BearerAuth
func (h *Handler) handleGetProductSpesification(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	data, err := h.store.GetProductSpesification(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if data == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	utils.WriteJson(w, http.StatusOK, data)
}

// @Summary delete data
// @Description delete data
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Router /admin/product/spesification/{id}/delete [delete]
// @Security BearerAuth
func (h *Handler) handleDeleteProductSpesification(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	data, err := h.store.GetProductSpesification(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if data == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	err = h.store.DeleteProductSpesification(data.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": "deleted"})
}

// @Summary update data
// @Summary update data
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param payload body types_product.ProductSpesificationPayload true "payload"
// @Router /admin/product/spesification/{id}/update [put]
// @Security BearerAuth
func (h *Handler) handleUpdateProductSpesification(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	data, err := h.store.GetProductSpesification(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if data == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	var payload types_product.ProductSpesificationPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, errors)
		return
	}

	product, err := h.store.GetProduct(payload.ProductId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if product == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}

	err = h.store.UpdateProductSpesification(data.Id, types_product.ProductSpesificationPayload{
		NameUz:    payload.NameUz,
		NameRu:    payload.NameRu,
		NameEn:    payload.NameEn,
		Brands:    payload.Brands,
		ProductId: payload.ProductId,
	})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "updated"})
}

// @Summary create product feature
// @Description create product feature
// @Tags product-admin
// @Accept json
// @Produce json
// @Param payload body types_product.ProductFeaturePayload true "payload"
// @Router /admin/product/feature/create [post]
// @Security BearerAuth
func (h *Handler) handleCreateProductFeature(w http.ResponseWriter, r *http.Request) {
	var payload types_product.ProductFeaturePayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	product, err := h.store.GetProduct(payload.ProductId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if product == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	err = h.store.CreateProductFeature(types_product.ProductFeaturePayload{
		TextUz:    payload.TextUz,
		TextRu:    payload.TextRu,
		TextEn:    payload.TextEn,
		ProductId: payload.ProductId,
	})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]string{"message": "created"})
}

// @Summary list product feature
// @Description list product feature
// @Tags product-admin
// @Accept json
// @Produce json
// @Router /admin/product/feature/list [get]
// @Security BearerAuth
func (h *Handler) handleListProductFeature(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.ListProductFeature()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, list)
}

// @Summary get feature
// @Description get feature
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Router /admin/product/feature/{id} [get]
// @Security BearerAuth
func (h *Handler) handleGetProductFeature(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	data, err := h.store.GetProductFeature(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if data == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	utils.WriteJson(w, http.StatusOK, data)
}

// @Summary delete feature
// @Description delete featrue
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Router /admin/product/feature/{id}/delete [delete]
// @Security BearerAuth
func (h *Handler) handleDeleteProductFeature(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	data, err := h.store.GetProductFeature(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if data == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	err = h.store.DeleteProductFeature(data.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": "deleted"})
}

// @Summary update feature
// @Description update feature
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param payload body types_product.ProductFeaturePayload true "payload"
// @Router /admin/product/feature/{id}/update [put]
// @Security BearerAuth
func (h *Handler) handleUpdateProductFeature(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	data, err := h.store.GetProductFeature(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if data == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("feature not found"))
		return
	}

	var payload types_product.ProductFeaturePayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	product, err := h.store.GetProduct(payload.ProductId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if product == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}

	err = h.store.UpdateProductFeature(data.Id, types_product.ProductFeaturePayload{
		TextUz:    payload.TextUz,
		TextRu:    payload.TextRu,
		TextEn:    payload.TextEn,
		ProductId: payload.ProductId,
	})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "updated"})
}

// @Summary create product feature
// @Description create product feature
// @Tags product-admin
// @Accept json
// @Produce json
// @Param payload body types_product.ProductAdventagePayload true "payload"
// @Router /admin/product/adventage/create [post]
// @Security BearerAuth
func (h *Handler) handleCreateProductAdventage(w http.ResponseWriter, r *http.Request) {
	var payload types_product.ProductAdventageListPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	product, err := h.store.GetProduct(payload.ProductId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if product == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	err = h.store.CreateProductAdvantage(types_product.ProductAdventagePayload{
		TextUz:    payload.TextUz,
		TextRu:    payload.TextRu,
		TextEn:    payload.TextEn,
		ProductId: payload.ProductId,
	})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]string{"message": "created"})
}

// @Summary list product feature
// @Description list product feature
// @Tags product-admin
// @Accept json
// @Produce json
// @Router /admin/product/adventage/list [get]
// @Security BearerAuth
func (h *Handler) handleListProductAdventage(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.ListProductAdventage()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, list)
}

// @Summary get feature
// @Description get feature
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Router /admin/product/adventage/{id} [get]
// @Security BearerAuth
func (h *Handler) handleGetProductAdventage(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	data, err := h.store.GetProductAdventage(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if data == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	utils.WriteJson(w, http.StatusOK, data)
}

// @Summary delete feature
// @Description delete featrue
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Router /admin/product/adventage/{id}/delete [delete]
// @Security BearerAuth
func (h *Handler) handleDeleteProductAdventage(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	data, err := h.store.GetProductAdventage(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if data == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	err = h.store.DeleteProductAdventage(data.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": "deleted"})
}

// @Summary update feature
// @Description update feature
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param payload body types_product.ProductAdventagePayload true "payload"
// @Router /admin/product/adventage/{id}/update [put]
// @Security BearerAuth
func (h *Handler) handleUpdateProductAdventage(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	data, err := h.store.GetProductAdventage(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if data == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("feature not found"))
		return
	}

	var payload types_product.ProductFeaturePayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	product, err := h.store.GetProduct(payload.ProductId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if product == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}

	err = h.store.UpdateProductAdventage(data.Id, types_product.ProductAdventagePayload{
		TextUz:    payload.TextUz,
		TextRu:    payload.TextRu,
		TextEn:    payload.TextEn,
		ProductId: payload.ProductId,
	})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "updated"})
}

// @Summary create product chemistry
// @Description create product chemistry
// @Tags product-admin
// @Accept json
// @Produce json
// @Param payload body types_product.ChemicalPropertyPayload true "payload"
// @Router /admin/product/chemistry/create [post]
// @Security BearerAuth
func (h *Handler) handleCreateChemistry(w http.ResponseWriter, r *http.Request) {
	var payload types_product.ChemicalPropertyPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	product, err := h.store.GetProduct(payload.ProductId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if product == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	err = h.store.CreateChemistry(types_product.ChemicalPropertyPayload{
		ProductId: payload.ProductId,
		NameUz:    payload.NameUz,
		NameRu:    payload.NameRu,
		NameEn:    payload.NameEn,
		Unit:      payload.Unit,
		Min:       payload.Min,
		Max:       payload.Max,
		Result:    payload.Result,
	})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]string{"message": "created"})
}

// @Summary list product chemistry
// @Description list product chemistry
// @Tags product-admin
// @Accept json
// @Produce json
// @Router /admin/product/chemistry/list [get]
// @Security BearerAuth
func (h *Handler) handleListChemistry(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.ListChemistry()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, list)
}

// @Summary get chemistry
// @Description get chemistry
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Router /admin/product/chemistry/{id} [get]
// @Security BearerAuth
func (h *Handler) handleGetChemistry(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	data, err := h.store.GetChemistry(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if data == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	utils.WriteJson(w, http.StatusOK, data)
}

// @Summary delete chemistry
// @Description delete chemistry
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Router /admin/product/chemistry/{id}/delete [delete]
// @Security BearerAuth
func (h *Handler) handleDeleteChemistry(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	data, err := h.store.GetChemistry(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if data == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	err = h.store.DeleteChemistry(data.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": "deleted"})
}

// @Summary update chemistry
// @Description update chemistry
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param payload body types_product.ChemicalPropertyPayload true "payload"
// @Router /admin/product/chemistry/{id}/update [put]
// @Security BearerAuth
func (h *Handler) handleUpdateChemistry(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	data, err := h.store.GetChemistry(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if data == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("feature not found"))
		return
	}

	var payload types_product.ChemicalPropertyPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	product, err := h.store.GetProduct(payload.ProductId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if product == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}

	err = h.store.UpdateChemistry(data.Id, types_product.ChemicalPropertyPayload{
		ProductId: payload.ProductId,
		NameUz:    payload.NameUz,
		NameRu:    payload.NameRu,
		NameEn:    payload.NameEn,
		Unit:      payload.Unit,
		Min:       payload.Min,
		Max:       payload.Max,
		Result:    payload.Result,
	})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "updated"})
}

// @Summary create product impact
// @Description create product impact
// @Tags product-admin
// @Accept json
// @Produce json
// @Param payload body types_product.ImpactPropertyPayload true "payload"
// @Router /admin/product/impact/create [post]
// @Security BearerAuth
func (h *Handler) handleCreateImpact(w http.ResponseWriter, r *http.Request) {
	var payload types_product.ImpactPropertyPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	product, err := h.store.GetProduct(payload.ProductId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if product == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	err = h.store.CreateImpact(types_product.ImpactPropertyPayload{
		ProductId:  payload.ProductId,
		MaterialUz: payload.MaterialUz,
		MaterialRu: payload.MaterialRu,
		MaterialEn: payload.MaterialEn,
		Unit:       payload.Unit,
		Max:        payload.Max,
		Result:     payload.Result,
	})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]string{"message": "created"})
}

// @Summary list product impact
// @Description list product impact
// @Tags product-admin
// @Accept json
// @Produce json
// @Router /admin/product/impact/list [get]
// @Security BearerAuth
func (h *Handler) handleListImpact(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.ListImpact()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, list)
}

// @Summary get impact
// @Description get impact
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Router /admin/product/impact/{id} [get]
// @Security BearerAuth
func (h *Handler) handleGetImpact(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	data, err := h.store.GetImpact(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if data == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	utils.WriteJson(w, http.StatusOK, data)
}

// @Summary delete impact
// @Description delete impact
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Router /admin/product/impact/{id}/delete [delete]
// @Security BearerAuth
func (h *Handler) handleDeleteImpact(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	data, err := h.store.GetImpact(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if data == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	err = h.store.DeleteImpact(data.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": "deleted"})
}

// @Summary update impact
// @Description update impact
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param payload body types_product.ImpactPropertyPayload true "payload"
// @Router /admin/product/impact/{id}/update [put]
// @Security BearerAuth
func (h *Handler) handleUpdateImpact(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	data, err := h.store.GetImpact(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if data == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("feature not found"))
		return
	}

	var payload types_product.ImpactPropertyPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	product, err := h.store.GetProduct(payload.ProductId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if product == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("product not found"))
		return
	}

	err = h.store.UpdateImpact(data.Id, types_product.ImpactPropertyPayload{
		ProductId:  payload.ProductId,
		MaterialUz: payload.MaterialUz,
		MaterialRu: payload.MaterialRu,
		MaterialEn: payload.MaterialEn,
		Unit:       payload.Unit,
		Max:        payload.Max,
		Result:     payload.Result,
	})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "updated"})
}

// @Summary create product file
// @Description create product file
// @Tags product-admin
// @Accept multipart/data
// @Produce json
// @Param productId formData string true "product id"
// @Param file formData file true "file"
// @Router /admin/product/file/create [post]
// @Security BearerAuth
func (h *Handler) handleCreateProductFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(50)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	productId := r.FormValue("productId")
	product, err := h.store.GetProduct(productId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if product == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	defer file.Close()

	filePath := "uploads/product_file/files/" + fileHeader.Filename

	outFile, err := os.Create(filePath)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	payload := types_product.ProductFilePayload{
		File:      filePath,
		ProductId: productId,
	}
	err = h.store.CreateProductFile(payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]string{"message": "created"})
}

// @Summary list product file
// @Description list product file
// @Tags product-admin
// @Accept json
// @Produce json
// @Router /admin/product/file/list [get]
// @Security BearerAuth
func (h *Handler) handleListProductFile(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.ListProductFile()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, list)
}

// @Summary get product file
// @Description get product file
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Router /admin/product/file/{id} [get]
// @Security BearerAuth
func (h *Handler) handleGetProductFile(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	data, err := h.store.GetProductFile(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if data == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	utils.WriteJson(w, http.StatusOK, data)
}

// @Summary delete product file
// @Description delete product file
// @Tags product-admin
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Router /admin/product/file/{id}/delete [delete]
// @Security BearerAuth
func (h *Handler) handleDeleteProductFile(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	data, err := h.store.GetProductFile(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if data == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	err = h.store.DeleteProductFile(data.Id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusNoContent, map[string]string{"message": "deleted"})
}

// @Summary update product file
// @Description update product file
// @Tags product-admin
// @Accept multipart/data
// @Produce json
// @Param id path string true "id"
// @Param productId formData string false "product id"
// @Param file formData file false "file"
// @Router /admin/product/file/{id}/update [put]
// @Security BearerAuth
func (h *Handler) handleUpdateProductFile(w http.ResponseWriter, r *http.Request) {
	var id = r.PathValue("id")
	data, err := h.store.GetProductFile(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if data == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
		return
	}

	err = r.ParseMultipartForm(50)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	productId := r.FormValue("productId")
	if productId != "" {
		product, err := h.store.GetProduct(productId)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
		if product == nil {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found"))
			return
		}
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		if err == http.ErrMissingFile {
			file = nil
			fileHeader = nil
		} else {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
	}

	if file != nil {
		defer file.Close()
	}

	var filePath string
	if fileHeader != nil {
		filePath = "uploads/product_file/files/" + fileHeader.Filename
	} else {
		filePath = ""
	}

	if filePath != "" {
		outFile, err := os.Create(filePath)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}

		defer outFile.Close()

		_, err = io.Copy(outFile, file)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
	}

	payload := types_product.ProductFilePayload{
		File:      filePath,
		ProductId: productId,
	}

	err = h.store.UpdateProductFile(id, payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]string{"message": "updated"})
}
