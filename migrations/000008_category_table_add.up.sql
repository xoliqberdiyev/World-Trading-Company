CREATE TABLE IF NOT EXISTS "categories"(
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "name_uz" VARCHAR NOT NULL,
    "name_ru" VARCHAR NOT NULL,
    "name_en" VARCHAR NOT NULL,
    "image" BYTEA NOT NULL,
    "icon" BYTEA NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);