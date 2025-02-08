CREATE TABLE IF NOT EXISTS "banner"(
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "image_uz" BYTEA NOT NULL,
    "image_ru" BYTEA NOT NULL,
    "image_en" BYTEA NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);