package common_admin

import (
	"database/sql"

	types_common_admin "github.com/XoliqberdiyevBehruz/wtc_backend/types/common_admin"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetContactUsList() ([]*types_common_admin.ContactListPayload, error) {
	var contacts []*types_common_admin.ContactListPayload
	query := `SELECT * FROM contact_us`
	rows, err := s.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	for rows.Next() {
		var contact types_common_admin.ContactListPayload
		if err := rows.Scan(&contact.Id, &contact.FirstName, &contact.LastName, &contact.Email, &contact.IsContacted, &contact.Comment, &contact.CreatedAt); err != nil {
			return nil, err
		}
		contacts = append(contacts, &contact)
	}
	return contacts, nil
}

func (s *Store) GetContactUsById(id string) (*types_common_admin.ContactListPayload, error) {
	var contact types_common_admin.ContactListPayload
	query := `SELECT * FROM contact_us WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&contact.Id, &contact.FirstName, &contact.LastName, &contact.Email, &contact.IsContacted, &contact.Comment, &contact.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &contact, err
}

func (s *Store) DeleteContactUsById(id string) error {
	query := `DELETE FROM contact_us WHERE id = $1`
	_, err := s.db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateContactById(id string, contact *types_common_admin.ContactUpdatePayload) error {
	query := `UPDATE contact_us SET is_contacted = $2 WHERE id = $1`
	_, err := s.db.Query(query, id, &contact.IsContacted)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CreateSettings(settings *types_common_admin.SettingsCreatePayload) error {
	query := `INSERT INTO settings(first_phone, second_phone, email, telegram_url, instagram_url, youtube_url, facebook_url, address_uz, address_ru, address_en, working_days) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := s.db.Exec(query, &settings.FirstPhone, &settings.SecondPhone, &settings.Email, &settings.Telegram, &settings.Instagram, &settings.Youtube, &settings.Facebook, &settings.AddressUz, &settings.AddressRu, &settings.AddressEn, &settings.WorkingDays)
	if err != nil {
		return err
	}
	return nil
}
// ========== settings ==========
func (s *Store) GetSettings() ([]*types_common_admin.SettingsPayload, error) {
	var settings []*types_common_admin.SettingsPayload
	query := `SELECT * FROM settings`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var setting types_common_admin.SettingsPayload
		if err := rows.Scan(&setting.Id, &setting.FirstPhone, &setting.SecondPhone, &setting.Email, &setting.Telegram, &setting.Instagram, &setting.Youtube, &setting.Facebook, &setting.AddressUz, &setting.AddressRu, &setting.AddressEn, &setting.WorkingDays, &setting.CreatedAt); err != nil {
			return nil, err
		}
		settings = append(settings, &setting)
	}
	return settings, nil
}

func (s *Store) GetSettingsById(id string) (*types_common_admin.SettingsPayload, error) {
	var settings types_common_admin.SettingsPayload
	query := `SELECT * FROM settings WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&settings.Id, &settings.FirstPhone, &settings.SecondPhone, &settings.Email, &settings.Telegram, &settings.Instagram, &settings.Youtube, &settings.Facebook, &settings.AddressUz, &settings.AddressRu, &settings.AddressEn,  &settings.WorkingDays, &settings.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &settings, nil
}

func (s *Store) UpdateSettingsById(id string, settings *types_common_admin.SettingsUpdatePayload) error {
	query := `
		UPDATE 
			settings 
		SET 
			first_phone = $1, second_phone = $2, email = $3, telegram_url = $4, 
			instagram_url = $5, youtube_url = $6, facebook_url = $7, address_uz = $9, address_ru = $10, address_en = $11, working_days = $12
		WHERE id = $8
	`
	_, err := s.db.Query(query, &settings.FirstPhone, &settings.SecondPhone, &settings.Email, &settings.Telegram, &settings.Instagram, &settings.Youtube, &settings.Facebook, id, &settings.AddressUz, &settings.AddressRu, &settings.AddressEn, &settings.WorkingDays)
	if err != nil {
		return err
	}
	return nil
}
