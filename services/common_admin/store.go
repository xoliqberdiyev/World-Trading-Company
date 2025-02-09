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
	err := s.db.QueryRow(query, id).Scan(&settings.Id, &settings.FirstPhone, &settings.SecondPhone, &settings.Email, &settings.Telegram, &settings.Instagram, &settings.Youtube, &settings.Facebook, &settings.AddressUz, &settings.AddressRu, &settings.AddressEn, &settings.WorkingDays, &settings.CreatedAt)
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

func (s *Store) GetAllContactUsFooter() ([]*types_common_admin.ContactUsFooterPayload, error) {
	var contacts []*types_common_admin.ContactUsFooterPayload
	query := `SELECT * FROM contact_us_footer`
	rows, err := s.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	for rows.Next() {
		var contactUs types_common_admin.ContactUsFooterPayload
		if err := rows.Scan(&contactUs.Id, &contactUs.FullName, &contactUs.Phone, &contactUs.Email, &contactUs.IsContacted, &contactUs.CreatedAt); err != nil {
			return nil, err
		}
		contacts = append(contacts, &contactUs)
	}
	return contacts, nil
}

func (s *Store) UpdateContactUsFooter(id string, contactUs *types_common_admin.ContactUsFooterUpdatePayload) error {
	query := `UPDATE contact_us_footer SET is_contacted = $1 WHERE id = $2`
	_, err := s.db.Query(query, &contactUs.IsContacted, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetContactUsFooterById(id string) (*types_common_admin.ContactUsFooterPayload, error) {
	var contactUs types_common_admin.ContactUsFooterPayload
	query := `SELECT * FROM contact_us_footer WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&contactUs.Id, &contactUs.FullName, &contactUs.Phone, &contactUs.Email, &contactUs.IsContacted, &contactUs.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &contactUs, err
}

func (s *Store) DeleteContactUsFooterById(id string) error {
	query := `DELETE FROM contact_us_footer WHERE id = $1`
	_, err := s.db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CreateMedia(media *types_common_admin.MediaPayload) error {
	query := `INSERT INTO medias(file_uz, file_ru, file_en, link) VALUES($1, $2, $3, $4)`
	_, err := s.db.Exec(query, &media.FileUz, &media.FileRu, &media.FileEn, &media.Link)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetMediaById(id string) (*types_common_admin.MediaListPayload, error) {
	var media types_common_admin.MediaListPayload
	query := `SELECT * FROM medias WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&media.Id, &media.FileUz, &media.FileRu, &media.FileEn, &media.Link, &media.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &media, err
}

func (s *Store) GetAllMedias() ([]*types_common_admin.MediaListPayload, error) {
	var medias []*types_common_admin.MediaListPayload
	query := `SELECT * FROM medias`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var media types_common_admin.MediaListPayload
		if err := rows.Scan(&media.Id, &media.FileUz, &media.FileRu, &media.FileEn, &media.Link, &media.CreatedAt); err != nil {
			return nil, err
		}
		medias = append(medias, &media)
	}
	return medias, nil
}

func (s *Store) DeleteMediaById(id string) error {
	query := `DELETE FROM medias WHERE id = $1`
	_, err := s.db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateMedia(id string, media *types_common_admin.MediaPayload) error {
	query := `UPDATE medias SET file_uz = $1, file_ru = $2, file_en = $3, link = $5 WHERE id = $4`
	_, err := s.db.Query(query, &media.FileUz, &media.FileRu, &media.FileEn, id, &media.Link)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CreatePartner(partren *types_common_admin.PartnersPayload) error {
	query := `INSERT INTO partners(image) VALUES($1)`
	_, err := s.db.Exec(query, &partren.Image)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdatePartner(id string, partner *types_common_admin.PartnersPayload) error {
	query := `UPDATE partners SET image = $1 WHERE id = $2`
	_, err := s.db.Query(query, &partner.Image, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) DeletePartner(id string) error {
	query := `DELETE FROM partners WHERE id = $1`
	_, err := s.db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetPartner(id string) (*types_common_admin.PartnersListPayload, error) {
	var partner types_common_admin.PartnersListPayload
	query := `SELECT * FROM partners WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&partner.Id, &partner.Image)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &partner, nil
}

func (s *Store) ListPartner() ([]*types_common_admin.PartnersListPayload, error) {
	var partners []*types_common_admin.PartnersListPayload
	query := `SELECT * FROM partners`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var partner types_common_admin.PartnersListPayload
		if err := rows.Scan(&partner.Id, &partner.Image); err != nil {
			return nil, err
		}
		partners = append(partners, &partner)
	}
	return partners, nil
}

func (s *Store) GetBanner(id string) (*types_common_admin.BannerListPayload, error) {
	var banner types_common_admin.BannerListPayload
	query := `SELECT * FROM banner WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&banner.Id, &banner.ImageUz, &banner.ImageRu, &banner.ImageEn, &banner.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &banner, nil
}

func (s *Store) CreateBanner(banner *types_common_admin.BannerPayload) (*types_common_admin.BannerListPayload, error) {
	var b types_common_admin.BannerListPayload
	query := `INSERT INTO banner(image_uz, image_ru, image_en) VALUES($1, $2, $3) RETURNING id, image_uz, image_ru, image_en, created_at`
	err := s.db.QueryRow(query, &banner.ImageUz, &banner.ImageRu, &banner.ImageEn).Scan(&b.Id, &b.ImageUz, &b.ImageRu, &b.ImageEn, &b.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (s *Store) UpdateBanner(id string, banner *types_common_admin.BannerPayload) (*types_common_admin.BannerListPayload, error) {
	var b types_common_admin.BannerListPayload
	query := `UPDATE banner SET image_uz = $1,  image_ru = $2, image_en = $3 WHERE id = $4 RETURNING id, image_uz, image_ru, image_en, created_at`
	err := s.db.QueryRow(query, &banner.ImageUz, &banner.ImageRu, &banner.ImageEn, id).Scan(&b.Id, &b.ImageUz, &b.ImageRu, &b.ImageEn, &b.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (s *Store) DeleteBanner(id string) error {
	query := `DELETE FROM banner WHERE id = $1`
	_, err := s.db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) ListBanner() ([]*types_common_admin.BannerListPayload, error) {
	var banners []*types_common_admin.BannerListPayload
	query := `SELECT * FROM banner`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b types_common_admin.BannerListPayload
		if err := rows.Scan(&b.Id, &b.ImageUz, &b.ImageRu, &b.ImageEn, &b.CreatedAt); err != nil {
			return nil, err
		}
		banners = append(banners, &b)
	}
	return banners, nil
}

func (s *Store) CreateNews(news *types_common_admin.NewsPayload) (*types_common_admin.NewsListPayload, error) {
	var n types_common_admin.NewsListPayload
	query := `
		INSERT INTO 
			news(title_uz, title_ru, title_en, description_uz, description_ru, description_en, image, link)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, title_uz, title_ru, title_en, description_uz, description_ru, description_en, image, link, created_at	
	`

	err := s.db.QueryRow(query, &news.TitleUz, &news.TitleRu, &news.TitleEn, &news.DescriptionUz, &news.DescriptionRu, &news.DescriptionEn, &news.Image, &news.Link).Scan(
		&n.Id, &n.TitleUz, &n.TitleRu, &n.TitleEn, &n.DescriptionUz, &n.DescriptionRu, &n.DescriptionEn, &n.Image, &n.Link, &n.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

func (s *Store) GetNews(id string) (*types_common_admin.NewsListPayload, error) {
	var news types_common_admin.NewsListPayload
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

func (s *Store) UpdateNews(id string, n *types_common_admin.NewsPayload) (*types_common_admin.NewsListPayload, error) {
	var news types_common_admin.NewsListPayload
	query := `
		UPDATE 
			news 
		SET 
			title_uz = $1, title_ru = $2, title_en = $3, 
			description_uz = $4, description_ru = $5, description_en = $6, 
			image = $7, link = $8
		WHERE 
			id = $9
		RETURNING
			id, title_uz, title_ru, title_en, description_uz, description_ru, description_en, image, link, created_at
	`
	err := s.db.QueryRow(query, &n.TitleUz, &n.TitleRu, &n.TitleEn, &n.DescriptionUz, &n.DescriptionRu, &n.DescriptionEn, &n.Image, &n.Link, id).Scan(
		&news.Id, &news.TitleUz, &news.TitleRu, &news.TitleEn, &news.DescriptionUz, &news.DescriptionRu, &news.DescriptionEn, &news.Image, &news.Link, &news.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &news, nil
}

func (s *Store) DeleteNews(id string) error {
	query := `DELETE FROM news WHERE id = $1`
	_, err := s.db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) ListNews(limit, offset int) ([]*types_common_admin.NewsListPayload, error) {
	var news []*types_common_admin.NewsListPayload
	query := `SELECT * FROM news ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var new types_common_admin.NewsListPayload
		if err := rows.Scan(&new.Id, &new.TitleUz, &new.TitleRu, &new.TitleEn, &new.DescriptionUz, &new.DescriptionRu, &new.DescriptionEn, &new.Image, &new.Link, &new.CreatedAt); err != nil {
			return nil, err
		}
		news = append(news, &new)
	}
	return news, nil
}

func (s *Store) CreateCertificate(payload *types_common_admin.CertificatePayload) (*types_common_admin.CertificateListPayload, error) {
	var certificate types_common_admin.CertificateListPayload
	query := `INSERT INTO certificates(name_uz, name_ru, name_en, text_uz, text_ru, text_en) VALUES($1, $2, $3, $4, $5, $6) RETURNING id, name_uz, name_ru, name_en, text_uz, text_ru, text_en, created_at`
	err := s.db.QueryRow(query, payload.NameUz, payload.NameRu, payload.NameEn, payload.TextUz, payload.TextRu, payload.TextEn).Scan(
		&certificate.Id, &certificate.NameUz, &certificate.NameRu, &certificate.NameEn, &certificate.TextUz, &certificate.TextRu, &certificate.TextEn, &certificate.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &certificate, nil
}

func (s *Store) GetCertificate(id string) (*types_common_admin.CertificateListPayload, error) {
	var certificate types_common_admin.CertificateListPayload
	query := `SELECT * FROM certificates WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&certificate.Id, &certificate.NameUz, &certificate.NameRu, &certificate.NameEn, &certificate.TextUz, &certificate.TextRu, &certificate.TextEn, &certificate.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &certificate, nil
}

func (s *Store) UpdateCertificate(id string, payload *types_common_admin.CertificatePayload) (*types_common_admin.CertificateListPayload, error) {
	var certificate types_common_admin.CertificateListPayload
	query := `UPDATE certificates SET name_uz = $1, name_ru = $2, name_en = $3, text_uz = $4, text_ru = $5, text_en = $6 WHERE id = $7 RETURNING id, name_uz, name_ru, name_en, text_uz, text_ru, text_en, created_at`
	err := s.db.QueryRow(query, payload.NameUz, payload.NameRu, payload.NameEn, payload.TextUz, payload.TextRu, payload.TextEn, id).Scan(
		&certificate.Id, &certificate.NameUz, &certificate.NameRu, &certificate.NameEn, &certificate.TextUz, &certificate.TextRu, &certificate.TextEn, &certificate.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &certificate, nil
}

func (s *Store) DeleteCertificate(id string) error {
	query := `DELETE FROM certificates WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) ListCertificate() ([]*types_common_admin.CertificateListPayload, error) {
	var certificates []*types_common_admin.CertificateListPayload
	query := `SELECT * FROM certificates ORDER BY created_at`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var certificate types_common_admin.CertificateListPayload
		if err := rows.Scan(&certificate.Id, &certificate.NameUz, &certificate.NameRu, &certificate.NameEn, &certificate.TextUz, &certificate.TextRu, &certificate.TextEn, &certificate.CreatedAt); err != nil {
			return nil, err
		}
		certificates = append(certificates, &certificate)
	}
	return certificates, nil
}
