package types_common_admin

import "time"

type CommonStore interface {
	GetContactUsList() ([]*ContactListPayload, error)
	GetContactUsById(id string) (*ContactListPayload, error)
	DeleteContactUsById(id string) error
	UpdateContactById(id string, contact *ContactUpdatePayload) error
	CreateSettings(settings *SettingsCreatePayload) error
	GetSettings() ([]*SettingsPayload, error)
	GetSettingsById(id string) (*SettingsPayload, error)
	UpdateSettingsById(id string, settings *SettingsUpdatePayload) error
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

// ========== settings ===========
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
