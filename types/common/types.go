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
	ListNews(limit, offset int) ([]*NewsListPayload, error)
	ListAboutOil() ([]*types_about_company.AboutOilListPayload, error)
	ListCertificate() ([]*CertificateListPayload, error)
	ListWhyUs() ([]*types_about_company.WhyUsListPayload, error)
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
	Id    string `json:"id"`
	Image string `json:"image"`
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
	CreatedAt time.Time `json:"createdAt"`
}
