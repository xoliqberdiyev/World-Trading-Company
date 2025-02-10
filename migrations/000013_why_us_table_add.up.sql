CREATE TABLE IF NOT EXISTS "why_us"(
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "title_uz" VARCHAR NOT NULL,
    "title_ru" VARCHAR NOT NULL,
    "title_en" VARCHAR NOT NULL,
    "description_uz" TEXT NOT NULL,
    "description_ru" TEXT NOT NULL,
    "description_en" TEXT NOT NULL,
    "image" BYTEA,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMp
);