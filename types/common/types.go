package types_common

type CommonStore interface {
	CreateContactUs(contact *ContactCreatePayload) error
	GetAllSettings() ([]*SettingsPayload, error)
	CreateContactUsFooter(contactUs ContactUsFooterPayload) error
	GetAllMedia() ([]*MediaPayload, error)
	ListPartner() ([]*PartnerListPayload, error)
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
	CreatedAt string `json:"createdAt"`
}

type PartnerListPayload struct {
	Id    string `json:"id"`
	Image string `json:"image"`
}
