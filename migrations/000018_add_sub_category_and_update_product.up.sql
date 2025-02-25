CREATE TABLE IF NOT EXISTS "sub_categories"(
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "name_uz" VARCHAR NOT NULL,
    "name_ru" VARCHAR NOT NULL,
    "name_en" VARCHAR NOT NULL,
    "icon" BYTEA NOT NULL,
    "category_id" UUID NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_category FOREIGN KEY ("category_id") REFERENCES "categories"("id") ON DELETE CASCADE
);

-- ALTER TABLE "products" 
-- ADD COLUMN "sub_category_id" UUID,
-- ADD CONSTRAINT fk_sub_category FOREIGN KEY ("sub_category_id") REFERENCES "sub_categories"("id");