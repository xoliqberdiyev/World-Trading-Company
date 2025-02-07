package types_about_company

type CompanyStore interface {
	CreateCapasity(payload *CapasityPayload) error
	GetCapasity(id string) (*CapasityListPayload, error)
	ListCapasity() ([]*CapasityListPayload, error)
	DeleteCapasity(id string) error
	UpdateCapasity(id string, capasity *CapasityPayload) error
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
