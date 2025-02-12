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

func (s *Store) CreateProduct(payload *types_product.ProductPayload) (*types_product.ProductListPayload, error) {
	var product types_product.ProductListPayload
	query := `
		INSERT INTO
			products(name_uz, name_ru, name_en, description_uz, description_ru, description_en, text_uz, text_ru, text_en, category_id, image, banner)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id, name_uz, name_ru, name_en, description_uz, description_ru, description_en, text_uz, text_ru, text_en, image, banner, created_at
	`
	err := s.db.QueryRow(
		query, payload.NameUz, payload.NameRu, payload.NameEn, payload.DescriptionUz, payload.DescriptionRu, payload.DescriptionEn,
		payload.TextUz, payload.TextRu, payload.TextEn, payload.CategoryId, payload.Image, payload.Banner,
	).Scan(
		&product.Id, &product.NameUz, &product.NameRu, &product.NameEn, &product.DescriptionUz, &product.DescriptionRu, &product.DescriptionEn,
		&product.TextUz, &product.TextRu, &product.TextEn, &product.Image, &payload.Banner, &product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *Store) ListProduct(offset, limit int) ([]*types_product.ProductListPayload, int, error) {
	var products []*types_product.ProductListPayload
	var count int
	countQuery := `SELECT COUNT(*) FROM products`
	err := s.db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		return nil, 0, err
	}
	query := `
		SELECT 
			id, name_uz, name_ru, name_en, description_uz, description_ru, description_en, text_uz, text_ru, text_en, image, banner, created_at
		FROM products
		ORDER BY created_at 
		DESC LIMIT $1 OFFSET $2
	`
	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0,  err
	}
	for rows.Next() {
		var product types_product.ProductListPayload
		err := rows.Scan(
			&product.Id, &product.NameUz, &product.NameRu, &product.NameEn, &product.DescriptionUz, &product.DescriptionRu,
			&product.DescriptionEn, &product.TextUz, &product.TextRu, &product.TextEn, &product.Image, &product.Banner, &product.CreatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		products = append(products, &product)
	}
	return products, count, nil
}

func (s *Store) GetProduct(id string) (*types_product.ProductListPayload, error) {
	var product types_product.ProductListPayload
	query := `SELECT id, name_uz, name_ru, name_en, description_uz, description_ru, description_en, text_uz, text_ru, text_en, image, banner, created_at FROM products WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(
		&product.Id, &product.NameUz, &product.NameRu, &product.NameEn, &product.DescriptionUz, &product.DescriptionRu, &product.DescriptionEn,
		&product.TextUz, &product.TextRu, &product.TextEn, &product.Image, &product.Banner, &product.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (s *Store) DeleteProduct(id string) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateProduct(id string, payload *types_product.ProductPayload) (*types_product.ProductListPayload, error) {
	var product types_product.ProductListPayload
	query := `UPDATE products SET `
	args := []interface{}{}
	index := 1

	if payload.NameUz != "" {
		query += fmt.Sprintf("name_uz = $%d, ", index)
		args = append(args, payload.NameUz)
		index++
	}
	if payload.NameRu != "" {
		query += fmt.Sprintf("name_ru = $%d, ", index)
		args = append(args, payload.NameRu)
		index++
	}
	if payload.NameEn != "" {
		query += fmt.Sprintf("name_en = $%d, ", index)
		args = append(args, payload.NameEn)
		index++
	}
	if payload.DescriptionUz != "" {
		query += fmt.Sprintf("description_uz = $%d, ", index)
		args = append(args, payload.DescriptionUz)
		index++
	}
	if payload.DescriptionRu != "" {
		query += fmt.Sprintf("description_ru = $%d, ", index)
		args = append(args, payload.DescriptionRu)
		index++
	}
	if payload.DescriptionEn != "" {
		query += fmt.Sprintf("description_en = $%d, ", index)
		args = append(args, payload.DescriptionEn)
		index++
	}
	if payload.TextUz != "" {
		query += fmt.Sprintf("text_uz = $%d, ", index)
		args = append(args, payload.TextUz)
		index++
	}
	if payload.TextRu != "" {
		query += fmt.Sprintf("text_ru = $%d, ", index)
		args = append(args, payload.TextRu)
		index++
	}
	if payload.TextEn != "" {
		query += fmt.Sprintf("text_en = $%d, ", index)
		args = append(args, payload.TextEn)
		index++
	}
	if payload.CategoryId != "" {
		query += fmt.Sprintf("category_id = $%d, ", index)
		args = append(args, payload.CategoryId)
		index++
	}
	if payload.Image != "" {
		query += fmt.Sprintf("image = $%d, ", index)
		args = append(args, payload.Image)
		index++
	}
	if payload.Banner != "" {
		query += fmt.Sprintf("banner = $%d, ", index)
		args = append(args, payload.Banner)
		index++
	}
	query = query[:len(query)-2] + fmt.Sprintf(" WHERE id = $%d RETURNING id, name_uz, name_ru, name_en, description_uz, description_ru, description_en, text_uz, text_ru, text_en, image, banner, created_at", index)
	args = append(args, id)
	err := s.db.QueryRow(query, args...).Scan(
		&product.Id, &product.NameUz, &product.NameRu, &product.NameEn, &product.DescriptionUz, &product.DescriptionRu, &product.DescriptionEn,
		&product.TextUz, &product.TextRu, &product.TextEn, &product.Image, &product.Banner, &product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *Store) CreateProductMedia(payload types_product.ProductMediaPayload) (*types_product.ProductMediaListPayload, error) {
	var productMedia types_product.ProductMediaListPayload
	query := `INSERT INTO product_medias(product_id, image) VALUES($1, $2) RETURNING id, image, created_at`
	err := s.db.QueryRow(query, payload.ProductId, payload.Image).Scan(&productMedia.Id, &productMedia.Image, &productMedia.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &productMedia, nil
}

func (s *Store) ListProductMedia(limit, offset int) ([]*types_product.ProductMediaListPayload, int, error) {
	var medias []*types_product.ProductMediaListPayload
	var count int
	countQuery := `SELECT COUNT(*) FROM product_medias`
	err := s.db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	query := `SELECT id, image, created_at FROM product_medias ORDER BY created_at DESC LIMIT $1 OFFSET $2	`
	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		var productMedia types_product.ProductMediaListPayload
		if err := rows.Scan(&productMedia.Id, &productMedia.Image, &productMedia.CreatedAt); err != nil {
			return nil, 0, err
		}
		medias = append(medias, &productMedia)
	}
	return medias, count, nil
}

func (s *Store) GetProductMedia(id string) (*types_product.ProductMediaListPayload, error) {
	var productMedia types_product.ProductMediaListPayload
	query := `SELECT id, image, created_at FROM product_medias WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&productMedia.Id, &productMedia.Image, &productMedia.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &productMedia, nil
}

func (s *Store) DeleteProductMedia(id string) error {
	query := `DELETE FROM product_medias WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateProductMedia(id string, payload types_product.ProductMediaPayload) (*types_product.ProductMediaListPayload, error) {
	var productMedia types_product.ProductMediaListPayload
	query := `UPDATE product_medias SET `
	args := []interface{}{}
	index := 1
	if payload.Image != "" {
		query += fmt.Sprintf("image = $%d, ", index)
		args = append(args, payload.Image)
		index++
	}
	if payload.ProductId != "" {
		query += fmt.Sprintf("product_id = $%d, ", index)
		args = append(args, payload.ProductId)
		index++
	}
	query = query[:len(query)-2] + fmt.Sprintf(" WHERE id = $%d RETURNING id, image, created_at", index)
	args = append(args, id)
	err := s.db.QueryRow(query, args...).Scan(&productMedia.Id, &productMedia.Image, &productMedia.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &productMedia, nil
}
