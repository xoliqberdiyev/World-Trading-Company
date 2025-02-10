package types_common_admin

import (
	"time"
)

type CommonStore interface {
	GetContactUsList() ([]*ContactListPayload, error)
	GetContactUsById(id string) (*ContactListPayload, error)
	DeleteContactUsById(id string) error
	UpdateContactById(id string, contact *ContactUpdatePayload) error
	CreateSettings(settings *SettingsCreatePayload) error
	GetSettings() ([]*SettingsPayload, error)
	GetSettingsById(id string) (*SettingsPayload, error)
	UpdateSettingsById(id string, settings *SettingsUpdatePayload) error
	GetAllContactUsFooter() ([]*ContactUsFooterPayload, error)
	UpdateContactUsFooter(id string, contactUs *ContactUsFooterUpdatePayload) error
	GetContactUsFooterById(id string) (*ContactUsFooterPayload, error)
	DeleteContactUsFooterById(id string) error
	CreateMedia(media *MediaPayload) error
	GetMediaById(id string) (*MediaListPayload, error)
	GetAllMedias() ([]*MediaListPayload, error)
	DeleteMediaById(id string) error
	UpdateMedia(id string, media *MediaPayload) error
	CreatePartner(partren *PartnersPayload) error
	UpdatePartner(id string, partner *PartnersPayload) error
	DeletePartner(id string) error
	GetPartner(id string) (*PartnersListPayload, error)
	ListPartner() ([]*PartnersListPayload, error)
	GetBanner(id string) (*BannerListPayload, error)
	CreateBanner(banner *BannerPayload) (*BannerListPayload, error)
	UpdateBanner(id string, banner *BannerPayload) (*BannerListPayload, error)
	DeleteBanner(id string) error
	ListBanner() ([]*BannerListPayload, error)
	CreateNews(news *NewsPayload) (*NewsListPayload, error)
	GetNews(id string) (*NewsListPayload, error)
	UpdateNews(id string, n *NewsPayload) (*NewsListPayload, error)
	DeleteNews(id string) error
	ListNews(limit, offset int) ([]*NewsListPayload, error)
	CreateCertificate(payload *CertificatePayload) (*CertificateListPayload, error)
	GetCertificate(id string) (*CertificateListPayload, error)
	UpdateCertificate(id string, payload *CertificatePayload) (*CertificateListPayload, error)
	DeleteCertificate(id string) error
	ListCertificate() ([]*CertificateListPayload, error)
}

type ContactListPayload struct {
	Id          string    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Comment     string    `json:"comment"`
	IsContacted bool      `json:"isContacted"`
	CreatedAt   time.Time `json:"createAt"`
}

type ContactUpdatePayload struct {
	IsContacted bool `json:"isContacted"`
}

type SettingsCreatePayload struct {
	FirstPhone  string `json:"firstPhone" validate:"required,e164"`
	SecondPhone string `json:"secondPhone" validate:"required,e164"`
	Email       string `json:"email" validate:"required,email"`
	Telegram    string `json:"telegramUrl"`
	Instagram   string `json:"instagramUrl"`
	Youtube     string `json:"youtubeUrl"`
	Facebook    string `json:"facebookUrl"`
	AddressUz   string `json:"addressUz"`
	AddressRu   string `json:"addressRu"`
	AddressEn   string `json:"addressEn"`
	WorkingDays string `json:"workingDays"`
}

type SettingsPayload struct {
	Id          string    `json:"id"`
	FirstPhone  string    `json:"firstPhone"`
	SecondPhone string    `json:"secondPhone"`
	Email       string    `json:"email"`
	Telegram    string    `json:"telegramUrl"`
	Instagram   string    `json:"instagramUrl"`
	Youtube     string    `json:"youtubeUrl"`
	Facebook    string    `json:"facebookUrl"`
	AddressUz   string    `json:"addressUz"`
	AddressRu   string    `json:"addressRu"`
	AddressEn   string    `json:"addressEn"`
	WorkingDays string    `json:"workingDays"`
	CreatedAt   time.Time `json:"createdAt"`
}

type SettingsUpdatePayload struct {
	FirstPhone  string `json:"firstPhone" validate:"required,e164"`
	SecondPhone string `json:"secondPhone" validate:"required,e164"`
	Email       string `json:"email" validate:"required,email"`
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
	Id          string    `json:"id"`
	FullName    string    `json:"fullName"`
	Phone       string    `json:"phoneNumber"`
	Email       string    `json:"email"`
	IsContacted bool      `json:"isContacted"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ContactUsFooterUpdatePayload struct {
	IsContacted bool `json:"isContacted"`
}

type MediaPayload struct {
	FileUz string `json:"fileUz"`
	FileRu string `json:"fileRu"`
	FileEn string `json:"fileEn"`
	Link   string `json:"link"`
}

type MediaListPayload struct {
	Id        string    `json:"id"`
	FileUz    string    `json:"fileUz"`
	FileRu    string    `json:"fileRu"`
	FileEn    string    `json:"fileEn"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"createdAt"`
}

type PartnersPayload struct {
	Image string `json:"image"`
}

type PartnersListPayload struct {
	Id    string `json:"id"`
	Image string `json:"image"`
}

type BannerPayload struct {
	ImageUz string `json:"imageUz"`
	ImageRu string `json:"imageRu"`
	ImageEn string `json:"imageEn"`
}

type BannerListPayload struct {
	Id        string    `json:"id"`
	ImageUz   string    `json:"imageUz"`
	ImageRu   string    `json:"imageRu"`
	ImageEn   string    `json:"imageEn"`
	CreatedAt time.Time `json:"createdAt"`
}

type NewsPayload struct {
	TitleUz       string `json:"titleUz"`
	TitleRu       string `json:"titleRu"`
	TitleEn       string `json:"titleEn"`
	DescriptionUz string `json:"descriptionUz"`
	DescriptionRu string `json:"descriptionRu"`
	DescriptionEn string `json:"descriptionEn"`
	Link          string `json:"link"`
	Image         string `json:"image"`
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

type CertificatePayload struct {
	NameUz string `json:"nameUz"`
	NameRu string `json:"nameRu"`
	NameEn string `json:"nameEn"`
	TextUz string `json:"textUz"`
	TextRu string `json:"textRu"`
	TextEn string `json:"textEn"`
	Image  string `json:"image"`
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
