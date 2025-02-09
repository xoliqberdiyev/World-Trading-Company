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
