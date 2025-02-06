CREATE TABLE IF NOT EXISTS "medias"(
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "file_uz" BYTEA NOT NULL,
    "file_ru" BYTEA,
    "file_en" BYTEA,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);