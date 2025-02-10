package product

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/XoliqberdiyevBehruz/wtc_backend/services/auth"
	types_product "github.com/XoliqberdiyevBehruz/wtc_backend/types/product"
	types_user_admin "github.com/XoliqberdiyevBehruz/wtc_backend/types/user_admin"
	"github.com/XoliqberdiyevBehruz/wtc_backend/utils"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store     types_product.ProductStore
	userStore types_user_admin.UserStore
}

func NewHandler(store types_product.ProductStore, userStore types_user_admin.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegsiterRoutes(r chi.Router) {
	r.Get("/admin/product/category/list", auth.AuthWithJWT(h.handleListCategory, h.userStore))
	r.Post("/admin/product/category/create", auth.AuthWithJWT(h.handleCreateCategory, h.userStore))
	r.Put("/admin/product/crategory/{categoryId}/update", auth.AuthWithJWT(h.handleUpdateCategory, h.userStore))
	r.Delete("/admin/product/category/{categoryId}/delete", auth.AuthWithJWT(h.handleDeleteCategory, h.userStore))
	r.Get("/admin/product/category/{categoryId}", auth.AuthWithJWT(h.handleGetCategory, h.userStore))
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
