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
	query := `INSERT INTO categories(name_uz, name_ru, name_en, icon) VALUES($1, $2, $3, $4)`
	_, err := s.db.Exec(query, &payload.NameUz, &payload.NameRu, &payload.NameEn, &payload.Icon)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetCategory(id string) (*types_product.CategoryListPayload, error) {
	var category types_product.CategoryListPayload
	query := `SELECT * FROM categories WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&category.Id, &category.NameUz, &category.NameRu, &category.NameEn, &category.Icon, &category.CreatedAt)
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
	query := `SELECT * FROM categories ORDER BY created_at`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category types_product.CategoryListPayload
		if err := rows.Scan(&category.Id, &category.NameUz, &category.NameRu, &category.NameEn, &category.Icon, &category.CreatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (s *Store) CreateProduct(payload *types_product.ProductPayload) (*types_product.ProductListPayload, error) {
	var product types_product.ProductListPayload
	var subCategoryId interface{}
	if payload.SubCategoryId == "" {
		subCategoryId = nil // Agar bo'sh bo'lsa, NULL yuboriladi
	} else {
		subCategoryId = payload.SubCategoryId // Agar qiymat bo'lsa, o'zini yuboradi
	}
	query := `
		INSERT INTO
			products(name_uz, name_ru, name_en, description_uz, description_ru, description_en, text_uz, text_ru, text_en, category_id, image, sub_category_id)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id, name_uz, name_ru, name_en, description_uz, description_ru, description_en, text_uz, text_ru, text_en, image, created_at
	`
	err := s.db.QueryRow(
		query, payload.NameUz, payload.NameRu, payload.NameEn, payload.DescriptionUz, payload.DescriptionRu, payload.DescriptionEn,
		payload.TextUz, payload.TextRu, payload.TextEn, payload.CategoryId, payload.Image, subCategoryId,
	).Scan(
		&product.Id, &product.NameUz, &product.NameRu, &product.NameEn, &product.DescriptionUz, &product.DescriptionRu, &product.DescriptionEn,
		&product.TextUz, &product.TextRu, &product.TextEn, &product.Image, &product.CreatedAt,
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
			id, name_uz, name_ru, name_en, description_uz, description_ru, description_en, text_uz, text_ru, text_en, image, created_at
		FROM products
		ORDER BY created_at 
		DESC LIMIT $1 OFFSET $2
	`
	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		var product types_product.ProductListPayload
		err := rows.Scan(
			&product.Id, &product.NameUz, &product.NameRu, &product.NameEn, &product.DescriptionUz, &product.DescriptionRu,
			&product.DescriptionEn, &product.TextUz, &product.TextRu, &product.TextEn, &product.Image, &product.CreatedAt,
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
	query := `SELECT id, name_uz, name_ru, name_en, description_uz, description_ru, description_en, text_uz, text_ru, text_en, image, created_at FROM products WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(
		&product.Id, &product.NameUz, &product.NameRu, &product.NameEn, &product.DescriptionUz, &product.DescriptionRu, &product.DescriptionEn,
		&product.TextUz, &product.TextRu, &product.TextEn, &product.Image, &product.CreatedAt,
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
	if payload.SubCategoryId != "" {
		query += fmt.Sprintf("sub_category_id = $%d, ", index)
		args = append(args, payload.SubCategoryId)
		index++
	}
	query = query[:len(query)-2] + fmt.Sprintf(" WHERE id = $%d RETURNING id, name_uz, name_ru, name_en, description_uz, description_ru, description_en, text_uz, text_ru, text_en, image, created_at", index)
	args = append(args, id)
	err := s.db.QueryRow(query, args...).Scan(
		&product.Id, &product.NameUz, &product.NameRu, &product.NameEn, &product.DescriptionUz, &product.DescriptionRu, &product.DescriptionEn,
		&product.TextUz, &product.TextRu, &product.TextEn, &product.Image, &product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *Store) CreateProductMedia(payload types_product.ProductMediaPayload) (*types_product.ProductMediaListPayload, error) {
	var productMedia types_product.ProductMediaListPayload
	query := `INSERT INTO product_medias(product_id, image, kilograms) VALUES($1, $2, $3) RETURNING id, image, created_at,, kilograms`
	err := s.db.QueryRow(query, payload.ProductId, payload.Image, payload.Kilograms).Scan(&productMedia.Id, &productMedia.Image, &productMedia.CreatedAt, &productMedia.Kilograms)
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

	query := `SELECT id, image, kilograms,  created_at FROM product_medias ORDER BY created_at DESC LIMIT $1 OFFSET $2	`
	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		var productMedia types_product.ProductMediaListPayload
		if err := rows.Scan(&productMedia.Id, &productMedia.Image, &productMedia.CreatedAt, &productMedia.Kilograms); err != nil {
			return nil, 0, err
		}
		medias = append(medias, &productMedia)
	}
	return medias, count, nil
}

func (s *Store) GetProductMedia(id string) (*types_product.ProductMediaListPayload, error) {
	var productMedia types_product.ProductMediaListPayload
	query := `SELECT id, image, kilograms, created_at FROM product_medias WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&productMedia.Id, &productMedia.Image, &productMedia.CreatedAt, &productMedia.Kilograms)
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
	if payload.Kilograms != "" {
		query += fmt.Sprintf("kilograms = $%d, ", index)
		args = append(args, payload.Kilograms)
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

func (s *Store) CreateProductSpesification(payload types_product.ProductSpesificationPayload) error {
	query := `
		INSERT INTO 
			product_specification(name_uz, name_ru, name_en, brands, product_id) 
		VALUES($1, $2, $3, $4, $5) 
	`
	_, err := s.db.Exec(query, payload.NameUz, payload.NameRu, payload.NameEn, payload.Brands, payload.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) ListProductSpesification() ([]*types_product.ProductListSpesificationPayload, error) {
	var list []*types_product.ProductListSpesificationPayload
	query := `SELECT * FROM product_specification`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var item types_product.ProductListSpesificationPayload
		if err := rows.Scan(&item.Id, &item.NameUz, &item.NameRu, &item.NameEn, &item.Brands, &item.ProductId, &item.CreatedAt); err != nil {
			return nil, err
		}
		list = append(list, &item)
	}
	return list, nil
}

func (s *Store) GetProductSpesification(id string) (*types_product.ProductListSpesificationPayload, error) {
	var data types_product.ProductListSpesificationPayload
	query := `SELECT * FROM product_specification WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&data.Id, &data.NameUz, &data.NameRu, &data.NameEn, &data.Brands, &data.ProductId, &data.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &data, nil
}

func (s *Store) DeleteProductSpesification(id string) error {
	query := `DELETE FROM product_specification WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateProductSpesification(id string, payload types_product.ProductSpesificationPayload) error {
	query := `UPDATE product_specification SET name_uz = $2, name_ru = $3, name_en = $4, brands = $5, product_id = $6 WHERE id = $1`
	_, err := s.db.Exec(query, id, payload.NameUz, payload.NameRu, payload.NameEn, payload.Brands, payload.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CreateProductFeature(payload types_product.ProductFeaturePayload) error {
	query := `INSERT INTO product_features(text_uz, text_ru, text_en, product_id) VALUES($1, $2, $3, $4)`
	_, err := s.db.Exec(query, payload.TextUz, payload.TextRu, payload.TextEn, payload.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) ListProductFeature() ([]*types_product.ProductFeatureListPayload, error) {
	var list []*types_product.ProductFeatureListPayload
	query := `SELECT * FROM product_features ORDER BY created_at`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var item types_product.ProductFeatureListPayload
		if err := rows.Scan(&item.Id, &item.TextUz, &item.TextRu, &item.TextEn, &item.ProductId, &item.CreatedAt); err != nil {
			return nil, err
		}
		list = append(list, &item)
	}
	return list, nil
}

func (s *Store) GetProductFeature(id string) (*types_product.ProductFeatureListPayload, error) {
	var feature types_product.ProductFeatureListPayload
	query := `SELECT * FROM product_features WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&feature.Id, &feature.TextUz, &feature.TextRu, &feature.TextEn, &feature.ProductId, &feature.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &feature, nil
}

func (s *Store) DeleteProductFeature(id string) error {
	query := `DELETE FROM product_features WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateProductFeature(id string, payload types_product.ProductFeaturePayload) error {
	query := `UPDATE product_features SET text_uz = $2, text_ru = $3, text_en = $4, product_id = $5 WHERE id = $1`
	_, err := s.db.Exec(query, id, payload.TextUz, payload.TextRu, payload.TextEn, payload.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CreateProductAdvantage(payload types_product.ProductAdventagePayload) error {
	query := `INSERT INTO product_adventage(text_uz, text_ru, text_en, product_id) VALUES($1, $2, $3, $4)`
	_, err := s.db.Exec(query, payload.TextUz, payload.TextRu, payload.TextEn, payload.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) ListProductAdventage() ([]*types_product.ProductAdventageListPayload, error) {
	var list []*types_product.ProductAdventageListPayload
	query := `SELECT * FROM product_adventage ORDER BY created_at`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var item types_product.ProductAdventageListPayload
		if err := rows.Scan(&item.Id, &item.TextUz, &item.TextRu, &item.TextEn, &item.ProductId, &item.CreatedAt); err != nil {
			return nil, err
		}
		list = append(list, &item)
	}
	return list, nil
}

func (s *Store) GetProductAdventage(id string) (*types_product.ProductAdventageListPayload, error) {
	var feature types_product.ProductAdventageListPayload
	query := `SELECT * FROM product_adventage WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&feature.Id, &feature.TextUz, &feature.TextRu, &feature.TextEn, &feature.ProductId, &feature.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &feature, nil
}

func (s *Store) DeleteProductAdventage(id string) error {
	query := `DELETE FROM product_adventage WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateProductAdventage(id string, payload types_product.ProductAdventagePayload) error {
	query := `UPDATE product_adventage SET text_uz = $2, text_ru = $3, text_en = $4, product_id = $5 WHERE id = $1`
	_, err := s.db.Exec(query, id, payload.TextUz, payload.TextRu, payload.TextEn, payload.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CreateChemistry(payload types_product.ChemicalPropertyPayload) error {
	query := `INSERT INTO chemical_property(product_id, name_uz, name_ru, name_en, unit, standard_range, analysis_result) VALUES($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.Exec(query, payload.ProductId, payload.NameUz, payload.NameRu, payload.NameEn, payload.Unit, payload.Range, payload.Result)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) ListChemistry() ([]*types_product.ChemicalPropertyListPayload, error) {
	var list []*types_product.ChemicalPropertyListPayload
	query := `SELECT * FROM chemical_property`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var item types_product.ChemicalPropertyListPayload
		if err := rows.Scan(&item.Id, &item.ProductId, &item.NameUz, &item.NameRu, &item.NameEn, &item.Unit, &item.Result, &item.Range); err != nil {
			return nil, err
		}
		list = append(list, &item)
	}
	return list, nil
}

func (s *Store) GetChemistry(id string) (*types_product.ChemicalPropertyListPayload, error) {
	var data types_product.ChemicalPropertyListPayload
	query := `SELECT * FROM chemical_property WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&data.Id, &data.ProductId, &data.NameUz, &data.NameRu, &data.NameEn, &data.Unit, &data.Result, &data.Range)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &data, nil
}

func (s *Store) DeleteChemistry(id string) error {
	query := `DELETE FROM chemical_property WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateChemistry(id string, payload types_product.ChemicalPropertyPayload) error {
	query := `UPDATE chemical_property SET product_id = $2, name_uz = $3, name_ru = $4, name_en = $5, unit = $6, analysis_result = $7, standard_range = $8 WHERE id = $1`
	_, err := s.db.Exec(query, id, payload.ProductId, payload.NameUz, payload.NameRu, payload.NameEn, payload.Unit, payload.Result, payload.Range)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CreateImpact(payload types_product.ImpactPropertyPayload) error {
	query := `INSERT INTO corrosion_impact(product_id, material_uz, material_ru, material_en, unit, max_limit, analysis_result) VALUES($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.Exec(query, payload.ProductId, payload.MaterialUz, payload.MaterialRu, payload.MaterialEn, payload.Unit, payload.Max, payload.Result)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) ListImpact() ([]*types_product.ImapctPropertyListPayload, error) {
	var list []*types_product.ImapctPropertyListPayload
	query := `SELECT * FROM corrosion_impact`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var item types_product.ImapctPropertyListPayload
		if err := rows.Scan(&item.Id, &item.ProductId, &item.MaterialUz, &item.MaterialRu, &item.MaterialEn, &item.Unit, &item.Max, &item.Result); err != nil {
			return nil, err
		}
		list = append(list, &item)
	}
	return list, nil
}

func (s *Store) GetImpact(id string) (*types_product.ImapctPropertyListPayload, error) {
	var data types_product.ImapctPropertyListPayload
	query := `SELECT * FROM corrosion_impact WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&data.Id, &data.ProductId, &data.MaterialUz, &data.MaterialRu, &data.MaterialEn, &data.Unit, &data.Max, &data.Result)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &data, nil
}

func (s *Store) DeleteImpact(id string) error {
	query := `DELETE FROM corrosion_impact WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateImpact(id string, payload types_product.ImpactPropertyPayload) error {
	query := `UPDATE corrosion_impact SET product_id = $2, material_uz = $3, material_ru = $4, material_en = $5, unit = $6, max_limit = $7, analysis_result = $8 WHERE id = $1`
	_, err := s.db.Exec(query, id, payload.ProductId, payload.MaterialUz, payload.MaterialRu, payload.MaterialEn, payload.Unit, payload.Max, payload.Result)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CreateProductFile(payload types_product.ProductFilePayload) error {
	query := `INSERT INTO product_files(file, product_id) VALUES($1, $2)`
	_, err := s.db.Exec(query, payload.File, payload.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) ListProductFile() ([]*types_product.ProductFileListPayload, error) {
	var list []*types_product.ProductFileListPayload
	query := `SELECT * FROM product_files`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item types_product.ProductFileListPayload
		if err := rows.Scan(&item.Id, &item.File, &item.ProductId, &item.CreatedAt); err != nil {
			return nil, err
		}
		list = append(list, &item)
	}
	return list, nil
}

func (s *Store) GetProductFile(id string) (*types_product.ProductFileListPayload, error) {
	var file types_product.ProductFileListPayload
	query := `SELECT * FROM product_files WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&file.Id, &file.File, &file.ProductId, &file.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &file, nil
}

func (s *Store) DeleteProductFile(id string) error {
	query := `DELETE FROM product_files WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateProductFile(id string, payload types_product.ProductFilePayload) error {
	query := `UPDATE product_files SET `
	args := []interface{}{}
	index := 1
	if payload.File != "" {
		query += fmt.Sprintf("file = $%d, ", index)
		args = append(args, payload.File)
		index++
	}
	if payload.ProductId != "" {
		query += fmt.Sprintf("product_id = $%d, ", index)
		args = append(args, payload.ProductId)
		index++
	}
	query = query[:len(query)-2] + fmt.Sprintf(" WHERE id = $%d", index)
	args = append(args, id)

	_, err := s.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) CreateSubCategory(payload types_product.SubCategroryPayload) error {
	query := `INSERT INTO sub_categories(name_uz, name_ru, name_en, icon, category_id) VALUES($1, $2, $3, $4, $5)`
	_, err := s.db.Exec(query, payload.NameUz, payload.NameRu, payload.NameEn, payload.Icon, payload.CategoryId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) ListSubCategory() ([]*types_product.SubCategoryListPayload, error) {
	var list []*types_product.SubCategoryListPayload
	query := `SELECT * FROM sub_categories`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var category types_product.SubCategoryListPayload
		if err := rows.Scan(&category.Id, &category.NameUz, &category.NameRu, &category.NameEn, &category.Icon, &category.CategoryId, &category.CreatedAt); err != nil {
			return nil, err
		}
		list = append(list, &category)
	}
	return list, nil
}

func (s *Store) GetSubCategory(id string) (*types_product.SubCategoryListPayload, error) {
	var category types_product.SubCategoryListPayload
	query := `SELECT * FROM sub_categories WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&category.Id, &category.NameUz, &category.NameRu, &category.NameEn, &category.Icon, &category.CategoryId, &category.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

func (s *Store) DeleteSubCategory(id string) error {
	query := `DELETE FROM sub_categories WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateSubCategory(id string, payload types_product.SubCategroryPayload) error {
	query := `UPDATE sub_categories SET `
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
	if payload.Icon != "" {
		query += fmt.Sprintf("icon = $%d, ", index)
		args = append(args, payload.Icon)
		index++
	}
	if payload.CategoryId != "" {
		query += fmt.Sprintf("category_id = $%d, ", index)
		args = append(args, payload.CategoryId)
		index++
	}

	query = query[:len(query)-2] + fmt.Sprintf(" WHERE id = $%d", index)
	args = append(args, id)
	_, err := s.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetProductsBySubCategoryId(id string) ([]*types_product.ProductListPayload, error) {
	var list []*types_product.ProductListPayload
	query := `SELECT id, name_uz, name_ru, name_en, description_uz, description_ru, description_en, text_uz, text_ru, text_en, image, created_at FROM products WHERE sub_category_id = $1`
	rows, err := s.db.Query(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	for rows.Next() {
		var product types_product.ProductListPayload
		if err := rows.Scan(&product.Id, &product.NameUz, &product.NameRu, &product.NameEn, &product.DescriptionUz, &product.DescriptionRu, &product.DescriptionEn, &product.TextUz, &product.TextRu, &product.TextEn, &product.Image, &product.CreatedAt); err != nil {
			return nil, err
		}
		list = append(list, &product)
	}

	return list, nil
}
