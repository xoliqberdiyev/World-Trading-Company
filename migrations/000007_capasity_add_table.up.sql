CREATE TABLE IF NOT EXISTS "capasity"(
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "name_uz" VARCHAR NOT NULL,
    "name_ru" VARCHAR NOT NULL,
    "name_en" VARCHAR NOT NULL,
    "quantity" INTEGER NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);