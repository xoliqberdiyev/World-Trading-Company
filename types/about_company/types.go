package types_about_company

import "time"

type CompanyStore interface {
	CreateCapasity(payload *CapasityPayload) error
	GetCapasity(id string) (*CapasityListPayload, error)
	ListCapasity() ([]*CapasityListPayload, error)
	DeleteCapasity(id string) error
	UpdateCapasity(id string, capasity *CapasityPayload) error
	CreateAboutOil(payload *AboutOilPayload) (*AboutOilListPayload, error)
	GetAboutOil(id string) (*AboutOilListPayload, error)
	UpdateAboutOil(id string, payload *AboutOilPayload) (*AboutOilListPayload, error)
	DeleteAboutOil(id string) error
	ListAboutOil() ([]*AboutOilListPayload, error)
	CreateWhyUs(payload *WhyUsPayload) (*WhyUsListPayload, error)
	ListWhyUs() ([]*WhyUsListPayload, error)
	DeleteWhyUs(id string) error
	UpdateWhyUs(id string, payload *WhyUsPayload) (*WhyUsListPayload, error)
	GetWhyUs(id string) (*WhyUsListPayload, error)
	CreateAboutUs(payload *AboutUsPayload) (*AboutUsListPayload, error)
	ListAboutUs() ([]*AboutUsListPayload, error)
	GetAboutUs(id string) (*AboutUsListPayload, error)
	DeleteAboutUs(id string) error
	UpdateAboutUs(id string, payload *AboutUsPayload) (*AboutUsListPayload, error)
}

type CapasityListPayload struct {
	Id        string `json:"id"`
	NameUz    string `json:"nameUz"`
	NameRu    string `json:"nameRu"`
	NameEn    string `json:"nameEn"`
	Quantity  int32  `json:"quantity"`
	CreatedAt string `json:"createdAt"`
}

type CapasityPayload struct {
	NameUz   string `json:"nameUz"`
	NameRu   string `json:"nameRu"`
	NameEn   string `json:"nameEn"`
	Quantity int32  `json:"quantity"`
}

type AboutOilPayload struct {
	NameUz string `json:"nameUz"`
	NameRu string `json:"nameRu"`
	NameEn string `json:"nameEn"`
	TextUz string `json:"textUz"`
	TextRu string `json:"textRu"`
	TextEn string `json:"textEn"`
}

type AboutOilListPayload struct {
	Id        string    `json:"id"`
	NameUz    string    `json:"nameUz"`
	NameRu    string    `json:"nameRu"`
	NameEn    string    `json:"nameEn"`
	TextUz    string    `json:"textUz"`
	TextRu    string    `json:"textRu"`
	TextEn    string    `json:"textEn"`
	CreatedAt time.Time `json:"createdAt"`
}

type WhyUsPayload struct {
	TitleUz       string `json:"titleUz"`
	TitleRu       string `json:"titleRu"`
	TitleEn       string `json:"titleEn"`
	DescriptionUz string `json:"descriptionUz"`
	DescriptionRu string `json:"descriptionRu"`
	DescriptionEn string `json:"descriptionEn"`
	Image         string `json:"image"`
}

type WhyUsListPayload struct {
	Id            string    `json:"id"`
	TitleUz       string    `json:"titleUz"`
	TitleRu       string    `json:"titleRu"`
	TitleEn       string    `json:"titleEn"`
	DescriptionUz string    `json:"descriptionUz"`
	DescriptionRu string    `json:"descriptionRu"`
	DescriptionEn string    `json:"descriptionEn"`
	Image         string    `json:"image"`
	CreatedAt     time.Time `json:"createdAt"`
}

type AboutUsPayload struct {
	TitleUz       string `json:"titleUz"`
	TitleRu       string `json:"titleRu"`
	TitleEn       string `json:"titleEn"`
	DescriptionUz string `json:"descriptionUz"`
	DescriptionRu string `json:"descriptionRu"`
	DescriptionEn string `json:"descriptionEn"`
	ImageUz       string `json:"imageUz"`
	ImageRu       string `json:"imageRu"`
	ImageEn       string `json:"imageEn"`
}

type AboutUsListPayload struct {
	Id            string    `json:"id"`
	TitleUz       string    `json:"titleUz"`
	TitleRu       string    `json:"titleRu"`
	TitleEn       string    `json:"titleEn"`
	DescriptionUz string    `json:"descriptionUz"`
	DescriptionRu string    `json:"descriptionRu"`
	DescriptionEn string    `json:"descriptionEn"`
	ImageUz       string    `json:"imageUz"`
	ImageRu       string    `json:"imageRu"`
	ImageEn       string    `json:"imageEn"`
	CreatedAt     time.Time `json:"createdAt"`
}
