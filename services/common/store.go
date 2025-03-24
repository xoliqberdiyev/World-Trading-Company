package common

import (
	"database/sql"

	types_about_company "github.com/XoliqberdiyevBehruz/wtc_backend/types/about_company"
	types_common "github.com/XoliqberdiyevBehruz/wtc_backend/types/common"
	types_product "github.com/XoliqberdiyevBehruz/wtc_backend/types/product"
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

func (s *Store) CreateContactUsFooter(contactUs types_common.ContactUsFooterPayload) error {
	query := `INSERT INTO contact_us_footer(full_name, phone_number, email) VALUES($1, $2, $3)`
	_, err := s.db.Exec(query, &contactUs.FullName, &contactUs.Phone, &contactUs.Email)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetAllMedia() ([]*types_common.MediaPayload, error) {
	var medias []*types_common.MediaPayload
	query := `SELECT * FROM medias`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var media types_common.MediaPayload
		if err := rows.Scan(&media.Id, &media.FileUz, &media.FileRu, &media.FileEn, &media.Link, &media.CreatedAt); err != nil {
			return nil, err
		}
		medias = append(medias, &media)
	}
	return medias, nil
}

func (s *Store) ListPartner() ([]*types_common.PartnerListPayload, error) {
	var partners []*types_common.PartnerListPayload

	query := `SELECT * FROM partners ORDER BY created_at`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var partner types_common.PartnerListPayload
		if err := rows.Scan(&partner.Id, &partner.Logo, &partner.Name, &partner.Flag, &partner.PartnerName, &partner.Email, &partner.PhoneNumber, &partner.CreatedAt); err != nil {
			return nil, err
		}
		partners = append(partners, &partner)
	}
	return partners, nil
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
		if err := rows.Scan(&category.Id, &category.NameUz, &category.NameRu, &category.NameEn, &category.Icon, &category.CreatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (s *Store) ListBanner() ([]*types_common.BannerPayload, error) {
	var banners []*types_common.BannerPayload
	query := `SELECT * FROM banner`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b types_common.BannerPayload
		if err := rows.Scan(&b.Id, &b.ImageUz, &b.ImageRu, &b.ImageEn, &b.CreatedAt); err != nil {
			return nil, err
		}
		banners = append(banners, &b)
	}
	return banners, nil
}

func (s *Store) ListNews(limit, offset int) ([]*types_common.NewsListPayload, int, error) {
	var news []*types_common.NewsListPayload
	var count int
	countQuery := `SELECT COUNT(*) FROM news`
	err := s.db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		return nil, 0, err
	}
	query := `SELECT * FROM news ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var new types_common.NewsListPayload
		if err := rows.Scan(&new.Id, &new.TitleUz, &new.TitleRu, &new.TitleEn, &new.DescriptionUz, &new.DescriptionRu, &new.DescriptionEn, &new.Image, &new.Link, &new.CreatedAt); err != nil {
			return nil, 0, err
		}
		news = append(news, &new)
	}
	return news, count, nil
}

func (s *Store) GetNews(id string) (*types_common.NewsListPayload, error) {
	var news types_common.NewsListPayload
	query := `SELECT * FROM news WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&news.Id, &news.TitleUz, &news.TitleRu, &news.TitleEn, &news.DescriptionUz, &news.DescriptionRu, &news.DescriptionEn, &news.Image, &news.Link, &news.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &news, nil
}

func (s *Store) ListAboutOil() ([]*types_about_company.AboutOilListPayload, error) {
	var oils []*types_about_company.AboutOilListPayload
	query := `SELECT * FROM about_oil ORDER BY created_at`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var oil types_about_company.AboutOilListPayload
		if err := rows.Scan(&oil.Id, &oil.NameUz, &oil.NameRu, &oil.NameEn, &oil.TextUz, &oil.TextRu, &oil.TextEn, &oil.CreatedAt); err != nil {
			return nil, err
		}
		oils = append(oils, &oil)
	}
	return oils, nil
}

func (s *Store) ListCertificate() ([]*types_common.CertificateListPayload, error) {
	var certificates []*types_common.CertificateListPayload
	query := `SELECT * FROM certificates ORDER BY created_at`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var certificate types_common.CertificateListPayload
		if err := rows.Scan(&certificate.Id, &certificate.NameUz, &certificate.NameRu, &certificate.NameEn, &certificate.TextUz, &certificate.TextRu, &certificate.TextEn, &certificate.Image, &certificate.CreatedAt); err != nil {
			return nil, err
		}
		certificates = append(certificates, &certificate)
	}
	return certificates, nil
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

func (s *Store) ListAboutUs() ([]*types_about_company.AboutUsListPayload, error) {
	var list []*types_about_company.AboutUsListPayload
	query := `SELECT * FROM about_us ORDER BY created_at`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var aboutUs types_about_company.AboutUsListPayload
		err := rows.Scan(
			&aboutUs.Id, &aboutUs.TitleUz, &aboutUs.TitleRu,
			&aboutUs.TitleEn, &aboutUs.DescriptionUz, &aboutUs.DescriptionRu,
			&aboutUs.DescriptionEn, &aboutUs.ImageUz, &aboutUs.ImageRu, &aboutUs.ImageEn,
			&aboutUs.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, &aboutUs)
	}
	return list, nil
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

func (s *Store) GetProductsByCategoryId(categoryId string) (*types_product.CategoryDetailPayload, error) {
	var category types_product.CategoryDetailPayload
	queryCategory := `SELECT * FROM categories WHERE id = $1`
	err := s.db.QueryRow(queryCategory, categoryId).Scan(&category.Id, &category.NameUz, &category.NameRu, &category.NameEn, &category.Icon, &category.CreatedAt)
	if err != nil {
		return nil, err
	}
	subCategoryCount := 0
	querySubCategoryCount := `SELECT COUNT(*) FROM sub_categories WHERE category_id = $1`
	err = s.db.QueryRow(querySubCategoryCount, category.Id).Scan(&subCategoryCount)
	if err != nil {
		return nil, err
	}
	if subCategoryCount != 0 {
		query := `SELECT id, name_uz, name_ru, name_en, icon, created_at FROM sub_categories WHERE category_id = $1`
		rows, err := s.db.Query(query, category.Id)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var sub types_product.SubCategoryDetailListPayload
			if err := rows.Scan(&sub.Id, &sub.NameUz, &sub.NameRu, &sub.NameEn, &sub.Icon, &sub.CreatedAt); err != nil {
				return nil, err
			}
			query := `SELECT id, name_uz, name_ru, name_en, image, created_at FROM products WHERE sub_category_id = $1`
			rows, err := s.db.Query(query, sub.Id)
			if err != nil {
				return nil, err
			}
			for rows.Next() {
				var product types_product.CategoryProductListPayload
				err := rows.Scan(&product.Id, &product.NameUz, &product.NameRu, &product.NameEn, &product.Image, &product.CreatedAt)
				if err != nil {
					return nil, err
				}
				sub.Products = append(sub.Products, product)
			}
			category.SubCategories = append(category.SubCategories, sub)
		}
	} else {
		queryProduct := `SELECT id, name_uz, name_ru, name_en, image, created_at FROM products WHERE category_id = $1`
		rows, err := s.db.Query(queryProduct, category.Id)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var product types_product.CategoryProductListPayload
			err := rows.Scan(&product.Id, &product.NameUz, &product.NameRu, &product.NameEn, &product.Image, &product.CreatedAt)
			if err != nil {
				return nil, err
			}
			category.Products = append(category.Products, product)
		}
	}

	return &category, nil
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

func (s *Store) GetProductById(id string) (*types_common.ProductDeatilPayload, error) {
	var product types_common.ProductDeatilPayload
	var subCategoryId sql.NullString
	queryProduct := `SELECT * FROM products WHERE id = $1`
	err := s.db.QueryRow(queryProduct, id).Scan(
		&product.Id, &product.NameUz, &product.NameRu, &product.NameEn,
		&product.DescriptionUz, &product.DescriptionRu, &product.DescriptionEn,
		&product.TextUz, &product.TextRu, &product.TextEn, &product.CategoryId,
		&product.Image, &product.CreatedAt, &subCategoryId,
	)
	if err != nil {
		return nil, err
	}
	if subCategoryId.Valid {
		product.SubCategoryId = subCategoryId.String
	} else {
		product.SubCategoryId = ""
	}


	queries := map[string]string{
		"media":         `SELECT * FROM product_medias WHERE product_id = $1`,
		"specification": `SELECT * FROM product_specification WHERE product_id = $1`,
		"feature":       `SELECT * FROM product_features WHERE product_id = $1`,
		"adventage":     `SELECT * FROM product_adventage WHERE product_id = $1`,
		"chemical":      `SELECT * FROM chemical_property WHERE product_id = $1`,
		"impact":        `SELECT * FROM corrosion_impact WHERE product_id = $1`,
		"files":         `SELECT * FROM product_files WHERE product_id = $1`,
	}

	for key, query := range queries {
		rows, err := s.db.Query(query, id)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		switch key {
		case "media":
			for rows.Next() {
				var media types_common.ProductMedia
				if err := rows.Scan(&media.Id, &media.Image, &media.ProductId, &media.CreatedAt, &media.Kilogram); err != nil {
					return nil, err
				}
				product.ProductMedias = append(product.ProductMedias, media)
			}
		case "specification":
			for rows.Next() {
				var spec types_common.ProductSpesification
				if err := rows.Scan(&spec.Id, &spec.NameUz, &spec.NameRu, &spec.NameEn, &spec.Brands, &spec.ProductId, &spec.CreatedAt); err != nil {
					return nil, err
				}
				product.ProductSpesification = append(product.ProductSpesification, spec)
			}
		case "feature":
			for rows.Next() {
				var feature types_common.ProductFeature
				if err := rows.Scan(&feature.Id, &feature.TextUz, &feature.TextRu, &feature.TextEn, &feature.ProductId, &feature.CreatedAt); err != nil {
					return nil, err
				}
				product.ProductFeature = append(product.ProductFeature, feature)
			}
		case "adventage":
			for rows.Next() {
				var adv types_common.ProductAdventage
				if err := rows.Scan(&adv.Id, &adv.TextUz, &adv.TextRu, &adv.TextEn, &adv.ProductId, &adv.CreatedAt); err != nil {
					return nil, err
				}
				product.ProductAdventage = append(product.ProductAdventage, adv)
			}
		case "chemical":
			for rows.Next() {
				var chem types_common.ChemicalProperty
				if err := rows.Scan(&chem.Id, &chem.ProductId, &chem.NameUz, &chem.NameRu, &chem.NameEn, &chem.Unit, &chem.Result, &chem.Range); err != nil {
					return nil, err
				}
				product.ChemicalProperty = append(product.ChemicalProperty, chem)
			}
		case "impact":
			for rows.Next() {
				var impact types_common.ImapctProperty
				if err := rows.Scan(&impact.Id, &impact.ProductId, &impact.MaterialUz, &impact.MaterialRu, &impact.MaterialEn, &impact.Unit, &impact.Max, &impact.Result); err != nil {
					return nil, err
				}
				product.ImapctProperty = append(product.ImapctProperty, impact)
			}
		case "files":
			for rows.Next() {
				var file types_common.ProductFile
				if err := rows.Scan(&file.Id, &file.File, &file.ProductId, &file.CreatedAt, &file.Kilogram); err != nil {
					return nil, err
				}
				product.ProductFile = append(product.ProductFile, file)
			}
		}
	}

	return &product, nil
}
