package common

import (
	"database/sql"

	types_common "github.com/XoliqberdiyevBehruz/wtc_backend/types/common"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateContactUs(contact *types_common.ContactCreatePayload) error {
	query := `INSERT INTO contact_us(first_name, last_name,  email, comment) VALUES($1, $2, $3, $4)`
	_, err := s.db.Exec(query, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Comment)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetAllSettings() ([]*types_common.SettingsPayload, error) {
	var settings []*types_common.SettingsPayload
	query := `SELECT first_phone, second_phone, email, telegram_url, instagram_url, youtube_url, facebook_url, address_uz, address_ru, address_en, working_days FROM settings`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var setting types_common.SettingsPayload
		if err := rows.Scan(&setting.FirstPhone, &setting.SecondPhone, &setting.Email, &setting.Telegram, &setting.Instagram, &setting.Youtube, &setting.Facebook, &setting.AddressUz, &setting.AddressRu, &setting.AddressEn, &setting.WorkingDays); err != nil {
			return nil, err
		}
		settings = append(settings, &setting)
	}
	return settings, nil
}
