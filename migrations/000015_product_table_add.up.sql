CREATE TABLE IF NOT EXISTS "products"(
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "name_uz" VARCHAR NOT NULL,
    "name_ru" VARCHAR NOT NULL,
    "name_en" VARCHAR NOT NULL,
    "description_uz" TEXT NOT NULL,
    "description_ru" TEXT NOT NULL,
    "description_en" TEXT NOT NULL,
    "text_uz" TEXT NOT NULL,
    "text_ru" TEXT NOT NULL,
    "text_en" TEXT NOT NULL,
    "category_id" UUID NOT NULL,
    "image" BYTEA,
    "banner" BYTEA,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_category FOREIGN KEY ("category_id") REFERENCES "categories"("id") ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "product_medias"(
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "image" BYTEA NOT NULL,
    "product_id" UUID NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_product FOREIGN KEY ("product_id") REFERENCES "products"("id") ON DELETE CASCADE
)