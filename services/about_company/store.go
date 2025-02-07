package about_company

import (
	"database/sql"

	types_about_company "github.com/XoliqberdiyevBehruz/wtc_backend/types/about_company"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateCapasity(payload *types_about_company.CapasityPayload) error {
	query := `INSERT INTO capasity(name_uz, name_ru, name_en, quantity) VALUES($1, $2, $3, $4)`
	_, err := s.db.Exec(query, &payload.NameUz, &payload.NameRu, &payload.NameEn, &payload.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetCapasity(id string) (*types_about_company.CapasityListPayload, error) {
	var capasity types_about_company.CapasityListPayload
	query := `SELECT * FROM capasity WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&capasity.Id, &capasity.NameUz, &capasity.NameRu, &capasity.NameEn, &capasity.Quantity, &capasity.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &capasity, nil
}

func (s *Store) ListCapasity() ([]*types_about_company.CapasityListPayload, error) {
	var capasities []*types_about_company.CapasityListPayload
	query := `SELECT * FROM capasity`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var capasity types_about_company.CapasityListPayload
		if err := rows.Scan(&capasity.Id, &capasity.NameUz, &capasity.NameRu, &capasity.NameEn, &capasity.Quantity, &capasity.CreatedAt); err != nil {
			return nil, err
		}
		capasities = append(capasities, &capasity)
	}
	return capasities, nil
}

func (s *Store) DeleteCapasity(id string) error {
	query := `DELETE FROM capasity WHERE id = $1`
	_, err := s.db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateCapasity(id string, capasity *types_about_company.CapasityPayload) error {
	query := `UPDATE capasity SET name_uz = $2, name_ru = $3, name_en = $4, quantity = $5 WHERE id = $1`
	_, err := s.db.Query(query, id, &capasity.NameUz, &capasity.NameRu, &capasity.NameEn, &capasity.Quantity)
	if err != nil {
		return err
	}
	return nil
} 
