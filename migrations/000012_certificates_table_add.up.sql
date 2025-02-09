CREATE TABLE IF NOT EXISTS "certificates" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "name_uz" VARCHAR NOT NULL,
    "name_ru" VARCHAR NOT NULL,
    "name_en" VARCHAR NOT NULL,
    "text_uz" TEXT NOT NULL,
    "text_ru" TEXT NOT NULL, 
    "text_en" TEXT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);