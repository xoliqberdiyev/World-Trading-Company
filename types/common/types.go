package types_common

import (
	"time"

	types_about_company "github.com/XoliqberdiyevBehruz/wtc_backend/types/about_company"
	types_product "github.com/XoliqberdiyevBehruz/wtc_backend/types/product"
)

type CommonStore interface {
	CreateContactUs(contact *ContactCreatePayload) error
	GetAllSettings() ([]*SettingsPayload, error)
	CreateContactUsFooter(contactUs ContactUsFooterPayload) error
	GetAllMedia() ([]*MediaPayload, error)
	ListPartner() ([]*PartnerListPayload, error)
	ListCategory() ([]*types_product.CategoryListPayload, error)
	ListBanner() ([]*BannerPayload, error)
	GetNews(id string) (*NewsListPayload, error)
	ListNews(limit, offset int) ([]*NewsListPayload, int, error)
	ListAboutOil() ([]*types_about_company.AboutOilListPayload, error)
	ListCertificate() ([]*CertificateListPayload, error)
	ListWhyUs() ([]*types_about_company.WhyUsListPayload, error)
	ListAboutUs() ([]*types_about_company.AboutUsListPayload, error)
	ListCapasity() ([]*types_about_company.CapasityListPayload, error)
	GetProductsByCategoryId(categoryId string) (*types_product.CategoryDetailPayload, error)
	GetCategory(id string) (*types_product.CategoryListPayload, error)
	GetProductById(id string) (*ProductDeatilPayload, error)
}

type ContactCreatePayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" validate:"required,email"`
	Comment   string `json:"comment"`
}

type SettingsPayload struct {
	FirstPhone  string `json:"firstPhone"`
	SecondPhone string `json:"secondPhone"`
	Email       string `json:"email"`
	Telegram    string `json:"telegramUrl"`
	Instagram   string `json:"instagramUrl"`
	Youtube     string `json:"youtubeUrl"`
	Facebook    string `json:"facebookUrl"`
	AddressUz   string `json:"addressUz"`
	AddressRu   string `json:"addressRu"`
	AddressEn   string `json:"addressEn"`
	WorkingDays string `json:"workingDays"`
}

type ContactUsFooterPayload struct {
	FullName string `json:"fullName"`
	Phone    string `json:"phone" validate:"required,e164"`
	Email    string `json:"email" validate:"required,email"`
}

type MediaPayload struct {
	Id        string `json:"id"`
	FileUz    string `json:"fileUz"`
	FileRu    string `json:"fileRu"`
	FileEn    string `json:"fileEn"`
	Link      string `json:"link"`
	CreatedAt string `json:"createdAt"`
}

type PartnerListPayload struct {
	Id          string    `json:"id"`
	Logo        string    `json:"logo"`
	Name        string    `json:"name"`
	Flag        string    `json:"flag"`
	PartnerName string    `json:"partnerName"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	CreatedAt   time.Time `json:"createdAt"`
}

type BannerPayload struct {
	Id        string    `json:"id"`
	ImageUz   string    `json:"imageUz"`
	ImageRu   string    `json:"ImageRu"`
	ImageEn   string    `json:"ImageEn"`
	CreatedAt time.Time `json:"createdAt"`
}

type NewsListPayload struct {
	Id            string    `json:"id"`
	TitleUz       string    `json:"titleUz"`
	TitleRu       string    `json:"titleRu"`
	TitleEn       string    `json:"titleEn"`
	DescriptionUz string    `json:"descriptionUz"`
	DescriptionRu string    `json:"descriptionRu"`
	DescriptionEn string    `json:"descriptionEn"`
	Link          string    `json:"link"`
	Image         string    `json:"image"`
	CreatedAt     time.Time `json:"createdAt"`
}

type CertificateListPayload struct {
	Id        string    `json:"id"`
	NameUz    string    `json:"nameUz"`
	NameRu    string    `json:"nameRu"`
	NameEn    string    `json:"nameEn"`
	TextUz    string    `json:"textUz"`
	TextRu    string    `json:"textRu"`
	TextEn    string    `json:"textEn"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"createdAt"`
}

type ProductDeatilPayload struct {
	Id                   string                 `json:"id"`
	NameUz               string                 `json:"nameUz"`
	NameRu               string                 `json:"nameRu"`
	NameEn               string                 `json:"nameEn"`
	DescriptionUz        string                 `json:"descriptionUz"`
	DescriptionRu        string                 `json:"descriptionRu"`
	DescriptionEn        string                 `json:"descriptionEn"`
	TextUz               string                 `json:"textUz"`
	TextRu               string                 `json:"textRu"`
	TextEn               string                 `json:"textEn"`
	Image                string                 `json:"image"`
	CreatedAt            time.Time              `json:"createdAt"`
	ProductMedias        []ProductMedia         `json:"productMedias"`
	ProductSpesification []ProductSpesification `json:"productSpesification"`
	ProductFeature       []ProductFeature       `json:"productFeature"`
	ProductAdventage     []ProductAdventage     `json:"productAdventage"`
	ChemicalProperty     []ChemicalProperty     `json:"productChemical"`
	ImapctProperty       []ImapctProperty       `json:"productImpact"`
	ProductFile          []ProductFile          `json:"productFile"`
}

type ProductMedia struct {
	Id        string    `json:"id"`
	Image     string    `json:"image"`
	ProductId string    `json:"productId"`
	CreatedAt time.Time `json:"createdAt"`
}

type ProductSpesification struct {
	Id        string    `json:"id"`
	NameUz    string    `json:"nameUz"`
	NameRu    string    `json:"nameRu"`
	NameEn    string    `json:"nameEn"`
	Brands    string    `json:"brands"`
	ProductId string    `json:"productId"`
	CreatedAt time.Time `json:"createdAt"`
}

type ProductFeature struct {
	Id        string    `json:"id"`
	TextUz    string    `json:"textUz"`
	TextRu    string    `json:"textRu"`
	TextEn    string    `json:"textEn"`
	ProductId string    `json:"productId"`
	CreatedAt time.Time `json:"createdAt"`
}

type ProductAdventage struct {
	Id        string    `json:"id"`
	TextUz    string    `json:"textUz"`
	TextRu    string    `json:"textRu"`
	TextEn    string    `json:"textEn"`
	ProductId string    `json:"productId"`
	CreatedAt time.Time `json:"createdAt"`
}

type ChemicalProperty struct {
	Id        string  `json:"id"`
	ProductId string  `json:"productId"`
	NameUz    string  `json:"nameUz"`
	NameRu    string  `json:"nameRu"`
	NameEn    string  `json:"nameEn"`
	Unit      string  `json:"unit"`
	Min       float32 `json:"min"`
	Max       float32 `json:"max"`
	Result    float32 `json:"result"`
}

type ImapctProperty struct {
	Id         string  `json:"id"`
	ProductId  string  `json:"productId"`
	MaterialUz string  `json:"materialUz"`
	MaterialRu string  `json:"materialRu"`
	MaterialEn string  `json:"materialEn"`
	Unit       string  `json:"unit"`
	Max        float32 `json:"max"`
	Result     float32 `json:"result"`
}

type ProductFile struct {
	Id        string    `json:"id"`
	File      string    `json:"file"`
	ProductId string    `json:"productId"`
	CreatedAt time.Time `json:"createdAt"`
}
