package types_product

import "time"

type ProductStore interface {
	CreateCategory(payload *CategoryPayload) error
	GetCategory(id string) (*CategoryListPayload, error)
	UpdateCategory(id string, payload *CategoryPayload) error
	DeleteCategory(id string) error
	ListCategory() ([]*CategoryListPayload, error)
	CreateProduct(payload *ProductPayload) (*ProductListPayload, error)
	ListProduct(offset, limit int) ([]*ProductListPayload, int, error)
	GetProduct(id string) (*ProductListPayload, error)
	DeleteProduct(id string) error
	UpdateProduct(id string, payload *ProductPayload) (*ProductListPayload, error)
	CreateProductMedia(payload ProductMediaPayload) (*ProductMediaListPayload, error)
	ListProductMedia(limit, offset int) ([]*ProductMediaListPayload, int, error)
	GetProductMedia(id string) (*ProductMediaListPayload, error)
	DeleteProductMedia(id string) error
	UpdateProductMedia(id string, payload ProductMediaPayload) (*ProductMediaListPayload, error)
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

type CategoryDetailPayload struct {
	Id        string                       `json:"id"`
	NameUz    string                       `json:"nameUz"`
	NameRu    string                       `json:"nameRu"`
	NameEn    string                       `json:"nameEn"`
	Image     string                       `json:"image"`
	Icon      string                       `json:"icon"`
	CreatedAt time.Time                    `json:"createdAt"`
	Products  []CategoryProductListPayload `json:"products"`
}

type CategoryProductListPayload struct {
	Id        string    `json:"id"`
	NameUz    string    `json:"nameUz"`
	NameRu    string    `json:"nameRu"`
	NameEn    string    `json:"nameEn"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"createdAt"`
}

type ProductPayload struct {
	NameUz        string `json:"nameUz"`
	NameRu        string `json:"nameRu"`
	NameEn        string `json:"nameEn"`
	DescriptionUz string `json:"descriptionUz"`
	DescriptionRu string `json:"descriptionRu"`
	DescriptionEn string `json:"descriptionEn"`
	TextUz        string `json:"textUz"`
	TextRu        string `json:"textRu"`
	TextEn        string `json:"textEn"`
	CategoryId    string `json:"categoryId"`
	Image         string `json:"image"`
	Banner        string `json:"banner"`
}

type ProductListPayload struct {
	Id            string    `json:"id"`
	NameUz        string    `json:"nameUz"`
	NameRu        string    `json:"nameRu"`
	NameEn        string    `json:"nameEn"`
	DescriptionUz string    `json:"descriptionUz"`
	DescriptionRu string    `json:"descriptionRu"`
	DescriptionEn string    `json:"descriptionEn"`
	TextUz        string    `json:"textUz"`
	TextRu        string    `json:"textRu"`
	TextEn        string    `json:"textEn"`
	Image         string    `json:"image"`
	Banner        string    `json:"banner"`
	CreatedAt     time.Time `json:"createdAt"`
}

type ProductMediaPayload struct {
	ProductId string `json:"productId"`
	Image     string `json:"image"`
}

type ProductMediaListPayload struct {
	Id        string    `json:"id"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"createdAt"`
}
