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
	CreateProductSpesification(payload ProductSpesificationPayload) error
	ListProductSpesification() ([]*ProductListSpesificationPayload, error)
	GetProductSpesification(id string) (*ProductListSpesificationPayload, error)
	DeleteProductSpesification(id string) error
	UpdateProductSpesification(id string, payload ProductSpesificationPayload) error
	CreateProductFeature(payload ProductFeaturePayload) error
	ListProductFeature() ([]*ProductFeatureListPayload, error)
	GetProductFeature(id string) (*ProductFeatureListPayload, error)
	DeleteProductFeature(id string) error
	UpdateProductFeature(id string, payload ProductFeaturePayload) error
	CreateProductAdvantage(payload ProductAdventagePayload) error
	ListProductAdventage() ([]*ProductAdventageListPayload, error)
	GetProductAdventage(id string) (*ProductAdventageListPayload, error)
	DeleteProductAdventage(id string) error
	UpdateProductAdventage(id string, payload ProductAdventagePayload) error
	CreateChemistry(payload ChemicalPropertyPayload) error
	ListChemistry() ([]*ChemicalPropertyListPayload, error)
	GetChemistry(id string) (*ChemicalPropertyListPayload, error)
	DeleteChemistry(id string) error
	UpdateChemistry(id string, payload ChemicalPropertyPayload) error
	CreateImpact(payload ImpactPropertyPayload) error
	ListImpact() ([]*ImapctPropertyListPayload, error)
	GetImpact(id string) (*ImapctPropertyListPayload, error)
	DeleteImpact(id string) error
	UpdateImpact(id string, payload ImpactPropertyPayload) error
	CreateProductFile(payload ProductFilePayload) error
	ListProductFile() ([]*ProductFileListPayload, error)
	GetProductFile(id string) (*ProductFileListPayload, error)
	DeleteProductFile(id string) error
	UpdateProductFile(id string, payload ProductFilePayload) error
	CreateSubCategory(payload SubCategroryPayload) error
	ListSubCategory() ([]*SubCategoryListPayload, error)
	GetSubCategory(id string) (*SubCategoryListPayload, error)
	DeleteSubCategory(id string) error
	UpdateSubCategory(id string, payload SubCategroryPayload) error
	GetProductsBySubCategoryId(id string) ([]*ProductListPayload, error)
}

type CategoryPayload struct {
	NameUz string `json:"nameUz"`
	NameRu string `json:"nameRu"`
	NameEn string `json:"nameEn"`
	Icon   string `json:"icon"`
}

type CategoryListPayload struct {
	Id        string    `json:"id"`
	NameUz    string    `json:"nameUz"`
	NameRu    string    `json:"nameRu"`
	NameEn    string    `json:"nameEn"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"createdAt"`
}

type SubCategoryDetailListPayload struct {
	Id        string                       `json:"id"`
	NameUz    string                       `json:"nameUz"`
	NameRu    string                       `json:"nameRu"`
	NameEn    string                       `json:"nameEn"`
	Icon      string                       `json:"icon"`
	CreatedAt time.Time                    `json:"createdAt"`
	Products  []CategoryProductListPayload `json:"products"`
}

type CategoryDetailPayload struct {
	Id            string                         `json:"id"`
	NameUz        string                         `json:"nameUz"`
	NameRu        string                         `json:"nameRu"`
	NameEn        string                         `json:"nameEn"`
	Icon          string                         `json:"icon"`
	CreatedAt     time.Time                      `json:"createdAt"`
	SubCategories []SubCategoryDetailListPayload `json:"subCategories"`
	Products      []CategoryProductListPayload   `json:"products"`
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
	SubCategoryId string `json:"subCategoryId"`
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

type ProductSpesificationPayload struct {
	NameUz    string `json:"nameUz"`
	NameRu    string `json:"nameRu"`
	NameEn    string `json:"nameEn"`
	Brands    string `json:"brands"`
	ProductId string `json:"productId"`
}

type ProductListSpesificationPayload struct {
	Id        string    `json:"id"`
	NameUz    string    `json:"nameUz"`
	NameRu    string    `json:"nameRu"`
	NameEn    string    `json:"nameEn"`
	Brands    string    `json:"brands"`
	ProductId string    `json:"productId"`
	CreatedAt time.Time `json:"createdAt"`
}

type ProductFeaturePayload struct {
	TextUz    string `json:"textUz"`
	TextRu    string `json:"textRu"`
	TextEn    string `json:"textEn"`
	ProductId string `json:"productId"`
}

type ProductFeatureListPayload struct {
	Id        string    `json:"id"`
	TextUz    string    `json:"textUz"`
	TextRu    string    `json:"textRu"`
	TextEn    string    `json:"textEn"`
	ProductId string    `json:"productId"`
	CreatedAt time.Time `json:"createdAt"`
}

type ProductAdventagePayload struct {
	TextUz    string `json:"textUz"`
	TextRu    string `json:"textRu"`
	TextEn    string `json:"textEn"`
	ProductId string `json:"productId"`
}

type ProductAdventageListPayload struct {
	Id        string    `json:"id"`
	TextUz    string    `json:"textUz"`
	TextRu    string    `json:"textRu"`
	TextEn    string    `json:"textEn"`
	ProductId string    `json:"productId"`
	CreatedAt time.Time `json:"createdAt"`
}

type ChemicalPropertyPayload struct {
	ProductId string  `json:"productId"`
	NameUz    string  `json:"nameUz"`
	NameRu    string  `json:"nameRu"`
	NameEn    string  `json:"nameEn"`
	Unit      string  `json:"unit"`
	Range     string  `json:"range"`
	Result    float32 `json:"result"`
}

type ChemicalPropertyListPayload struct {
	Id        string  `json:"id"`
	ProductId string  `json:"productId"`
	NameUz    string  `json:"nameUz"`
	NameRu    string  `json:"nameRu"`
	NameEn    string  `json:"nameEn"`
	Unit      string  `json:"unit"`
	Range     string  `json:"range"`
	Result    float32 `json:"result"`
}

type ImpactPropertyPayload struct {
	ProductId  string  `json:"productId"`
	MaterialUz string  `json:"materialUz"`
	MaterialRu string  `json:"materialRu"`
	MaterialEn string  `json:"materialEn"`
	Unit       string  `json:"unit"`
	Max        string  `json:"max"`
	Result     float32 `json:"result"`
}

type ImapctPropertyListPayload struct {
	Id         string  `json:"id"`
	ProductId  string  `json:"productId"`
	MaterialUz string  `json:"materialUz"`
	MaterialRu string  `json:"materialRu"`
	MaterialEn string  `json:"materialEn"`
	Unit       string  `json:"unit"`
	Max        string  `json:"max"`
	Result     float32 `json:"result"`
}

type ProductFilePayload struct {
	File      string `json:"file"`
	ProductId string `json:"productId"`
}

type ProductFileListPayload struct {
	Id        string    `json:"id"`
	File      string    `json:"file"`
	ProductId string    `json:"productId"`
	CreatedAt time.Time `json:"createdAt"`
}

type SubCategroryPayload struct {
	NameUz     string `json:"nameUz"`
	NameRu     string `json:"nameRu"`
	NameEn     string `json:"nameEn"`
	Icon       string `json:"icon"`
	CategoryId string `json:"categoryId"`
}

type SubCategoryListPayload struct {
	Id         string    `json:"id"`
	NameUz     string    `json:"nameUz"`
	NameRu     string    `json:"nameRu"`
	NameEn     string    `json:"nameEn"`
	Icon       string    `json:"icon"`
	CategoryId string    `json:"categoryId"`
	CreatedAt  time.Time `json:"createdAt"`
}
