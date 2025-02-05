CREATE TABLE IF NOT EXISTS "settings" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "first_phone" VARCHAR,
    "second_phone" VARCHAR,
    "email" VARCHAR NOT NULL,
    "telegram_url" VARCHAR NOT NULL,
    "instagram_url" VARCHAR NOT NULL,
    "youtube_url"  VARCHAR NOT NULL,
    "facebook_url" VARCHAR NOT NULL,
    "address_uz" VARCHAR NOT NULL,
    "address_ru" VARCHAR NOT NULL,
    "address_en" VARCHAR NOT NULL,
    "working_days" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);