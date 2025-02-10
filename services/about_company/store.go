package about_company

import (
	"database/sql"
	"fmt"

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

func (s *Store) CreateAboutOil(payload *types_about_company.AboutOilPayload) (*types_about_company.AboutOilListPayload, error) {
	var oil types_about_company.AboutOilListPayload
	query := `INSERT INTO about_oil(name_uz, name_ru, name_en, text_uz, text_ru, text_en) VALUES($1, $2, $3, $4, $5, $6) RETURNING id, name_uz, name_ru, name_en, text_uz, text_ru, text_en, created_at`
	err := s.db.QueryRow(query, payload.NameUz, payload.NameRu, payload.NameEn, payload.TextUz, payload.TextRu, payload.TextEn).Scan(&oil.Id, &oil.NameUz, &oil.NameRu, &oil.NameEn, &oil.TextUz, &oil.TextRu, &oil.TextEn, &oil.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &oil, nil
}

func (s *Store) GetAboutOil(id string) (*types_about_company.AboutOilListPayload, error) {
	var oil types_about_company.AboutOilListPayload
	query := `SELECT * FROM about_oil WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&oil.Id, &oil.NameUz, &oil.NameRu, &oil.NameEn, &oil.TextUz, &oil.TextRu, &oil.TextEn, &oil.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &oil, nil
}

func (s *Store) UpdateAboutOil(id string, payload *types_about_company.AboutOilPayload) (*types_about_company.AboutOilListPayload, error) {
	var oil types_about_company.AboutOilListPayload
	query := `UPDATE about_oil SET name_uz = $1, name_ru = $2, name_en = $3, text_uz = $4, text_ru = $5, text_en = $6 WHERE id = $7 RETURNING id, name_uz, name_ru, name_en, text_uz, text_ru, text_en, created_at`
	err := s.db.QueryRow(query, &payload.NameUz, &payload.NameRu, &payload.NameEn, &payload.TextUz, &payload.TextRu, &payload.TextEn, id).Scan(
		&oil.Id, &oil.NameUz, &oil.NameRu, &oil.NameEn, &oil.TextUz, &oil.TextRu, &oil.TextEn, &oil.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &oil, nil
}

func (s *Store) DeleteAboutOil(id string) error {
	query := `DELETE FROM about_oil WHERE id = $1`
	_, err := s.db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) ListAboutOil() ([]*types_about_company.AboutOilListPayload, error) {
	var oil []*types_about_company.AboutOilListPayload
	query := `SELECT * FROM about_oil ORDER BY created_at`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var o types_about_company.AboutOilListPayload
		if err := rows.Scan(&o.Id, &o.NameUz, &o.NameRu, &o.NameEn, &o.TextUz, &o.TextRu, &o.TextEn, &o.CreatedAt); err != nil {
			return nil, err
		}
		oil = append(oil, &o)
	}
	return oil, nil
}

func (s *Store) CreateWhyUs(payload *types_about_company.WhyUsPayload) (*types_about_company.WhyUsListPayload, error) {
	var whyUs types_about_company.WhyUsListPayload
	query := `INSERT INTO why_us(title_uz, title_ru, title_en, description_uz, description_ru, description_en, image) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id, title_uz, title_ru, title_en, description_uz, description_ru, description_en, image, created_at`
	err := s.db.QueryRow(query, payload.TitleUz, payload.TitleRu, payload.TitleEn, payload.DescriptionUz, payload.DescriptionRu, payload.DescriptionEn, payload.Image).Scan(
		&whyUs.Id, &whyUs.TitleUz, &whyUs.TitleRu, &whyUs.TitleEn, &whyUs.DescriptionUz, &whyUs.DescriptionRu, &whyUs.DescriptionEn, &whyUs.Image, &whyUs.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &whyUs, nil
}

func (s *Store) ListWhyUs() ([]*types_about_company.WhyUsListPayload, error) {
	var whyUsList []*types_about_company.WhyUsListPayload
	query := `SELECT * FROM why_us ORDER BY created_at`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var whyUs types_about_company.WhyUsListPayload
		if err := rows.Scan(&whyUs.Id, &whyUs.TitleUz, &whyUs.TitleRu, &whyUs.TitleEn, &whyUs.DescriptionUz, &whyUs.DescriptionRu, &whyUs.DescriptionEn, &whyUs.Image, &whyUs.CreatedAt); err != nil {
			return nil, err
		}
		whyUsList = append(whyUsList, &whyUs)
	}
	return whyUsList, nil
}

func (s *Store) DeleteWhyUs(id string) error {
	query := `DELETE FROM why_us WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateWhyUs(id string, payload *types_about_company.WhyUsPayload) (*types_about_company.WhyUsListPayload, error) {
	var whyUs types_about_company.WhyUsListPayload
	query := `UPDATE why_us SET `
	args := []interface{}{}
	argsIndex := 1

	if payload.TitleUz != "" {
		query += fmt.Sprintf("title_uz = $%d, ", argsIndex)
		args = append(args, payload.TitleUz)
		argsIndex++
	}


	if payload.TitleRu != "" {
		query += fmt.Sprintf("title_ru = $%d, ", argsIndex)
		args = append(args, payload.TitleRu)
		argsIndex++
	}

	if payload.TitleEn != "" {
		query += fmt.Sprintf("title_en = $%d, ", argsIndex)
		args = append(args, payload.TitleEn)
		argsIndex++
	}

	if payload.DescriptionUz != "" {
		query += fmt.Sprintf("description_uz = $%d, ", argsIndex)
		args = append(args, payload.DescriptionUz)
		argsIndex++
	}

	if payload.DescriptionRu != "" {
		query += fmt.Sprintf("description_ru = $%d, ", argsIndex)
		args = append(args, payload.DescriptionRu)
		argsIndex++
	}

	if payload.DescriptionEn != "" {
		query += fmt.Sprintf("description_en = $%d, ", argsIndex)
		args = append(args, payload.DescriptionEn)
		argsIndex++
	}

	if payload.Image != "" {
		query += fmt.Sprintf("image = $%d, ", argsIndex)
		args = append(args, payload.Image)
		argsIndex++
	}

	query = query[:len(query)-2] + fmt.Sprintf(" WHERE id = $%d RETURNING id, title_uz, title_ru, title_en, description_uz, description_ru, description_en, image, created_at", argsIndex)
	args = append(args, id)

	err := s.db.QueryRow(query, args...).Scan(
		&whyUs.Id, &whyUs.TitleUz, &whyUs.TitleRu, &whyUs.TitleEn, &whyUs.DescriptionUz, &whyUs.DescriptionRu, &whyUs.DescriptionEn, &whyUs.Image, &whyUs.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &whyUs, nil
}

func (s *Store) GetWhyUs(id string) (*types_about_company.WhyUsListPayload, error) {
	var whyUs types_about_company.WhyUsListPayload
	query := `SELECT * FROM why_us WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(
		&whyUs.Id, &whyUs.TitleUz, &whyUs.TitleRu, &whyUs.TitleEn, &whyUs.DescriptionUz, &whyUs.DescriptionRu, &whyUs.DescriptionEn, &whyUs.Image, &whyUs.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &whyUs, nil
}
