package types_product

import "time"

type ProductStore interface {
	CreateCategory(payload *CategoryPayload) error
	GetCategory(id string) (*CategoryListPayload, error)
	UpdateCategory(id string, payload *CategoryPayload) error
	DeleteCategory(id string) error
	ListCategory() ([]*CategoryListPayload, error)
}

type CategoryPayload struct {
	NameUz string `json:"nameUz"`
	NameRu string `json:"nameRu"`
	NameEn string `json:"nameEn"`
	Image  string `json:"image"`
	Icon   string `json:"icon"`
}

type CategoryListPayload struct {
	Id        string    `json:"id"`
	NameUz    string    `json:"nameUz"`
	NameRu    string    `json:"nameRu"`
	NameEn    string    `json:"nameEn"`
	Image     string    `json:"image"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"createdAt"`
}
