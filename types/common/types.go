package types_common

type CommonStore interface {
	CreateContactUs(contact *ContactCreatePayload) error
	GetAllSettings() ([]*SettingsPayload, error)
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
