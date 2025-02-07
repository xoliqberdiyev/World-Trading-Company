package product

import (
	"net/http"

	"github.com/XoliqberdiyevBehruz/wtc_backend/services/auth"
	types_product "github.com/XoliqberdiyevBehruz/wtc_backend/types/product"
	types_user_admin "github.com/XoliqberdiyevBehruz/wtc_backend/types/user_admin"
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

func (h *Handler) handleCreateCategory(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleListCategory(w http.ResponseWriter, r *http.Request) {
	
}

func (h *Handler) handleUpdateCategory(w http.ResponseWriter, r *http.Request) {
	
}

func (h *Handler) handleDeleteCategory(w http.ResponseWriter, r *http.Request) {
	
}

func (h *Handler) handleGetCategory(w http.ResponseWriter, r *http.Request) {
	
}
