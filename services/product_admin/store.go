package product

import (
	"database/sql"
	"fmt"

	types_product "github.com/XoliqberdiyevBehruz/wtc_backend/types/product"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateCategory(payload *types_product.CategoryPayload) error {
	query := `INSERT INTO categories(name_uz, name_ru, name_en, image, icon) VALUES($1, $2, $3, $4, $5)`
	_, err := s.db.Exec(query, &payload.NameUz, &payload.NameRu, &payload.NameEn, &payload.Image, &payload.Icon)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetCategory(id string) (*types_product.CategoryListPayload, error) {
	var category types_product.CategoryListPayload
	query := `SELECT * FROM categories WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&category.Id, &category.NameUz, &category.NameRu, &category.NameEn, &category.Image, &category.Icon, &category.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

func (s *Store) UpdateCategory(id string, payload *types_product.CategoryPayload) error {
	query := `UPDATE categories SET `
	args := []interface{}{}
	argsIndex := 1
	if payload.NameUz != "" {
		query += fmt.Sprintf("name_uz = $%d, ", argsIndex)
		args = append(args, payload.NameUz)
		argsIndex++
	}

	if payload.NameRu != "" {
		query += fmt.Sprintf("name_ru = $%d, ", argsIndex)
		args = append(args, payload.NameRu)
		argsIndex++
	}

	if payload.NameEn != "" {
		query += fmt.Sprintf("name_en = $%d, ", argsIndex)
		args = append(args, payload.NameEn)
		argsIndex++
	}

	if payload.Image != "" {
		query += fmt.Sprintf("image = $%d, ", argsIndex)
		args = append(args, payload.Image)
		argsIndex++
	}

	if payload.Icon != "" {
		query += fmt.Sprintf("icon = $%d, ", argsIndex)
		args = append(args, payload.Icon)
		argsIndex++
	}
	query = query[:len(query)-2] + fmt.Sprintf(" WHERE id = $%d", argsIndex)
	args = append(args, id)
	_, err := s.db.Query(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteCategory(id string) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := s.db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) ListCategory() ([]*types_product.CategoryListPayload, error) {
	var categories []*types_product.CategoryListPayload
	query := `SELECT * FROM categories`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category types_product.CategoryListPayload
		if err := rows.Scan(&category.Id, &category.NameUz, &category.NameRu, &category.NameEn, &category.Image, &category.Icon, &category.CreatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}
