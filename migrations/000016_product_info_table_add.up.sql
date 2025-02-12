CREATE TABLE IF NOT EXISTS "product_specification"(
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "name_uz" VARCHAR NOT NULL,
    "name_ru" VARCHAR NOT NULL,
    "name_en" VARCHAR NOT NULL,
    "brands" TEXT NOT NULL,
    "product_id" UUID NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_product FOREIGN KEY ("product_id") REFERENCES "products"("id") ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "product_features"(
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "text_uz" TEXT NOT NULL,
    "text_ru" TEXT NOT NULL,
    "text_en" TEXT NOT NULL,
    "product_id" UUID NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_product FOREIGN KEY ("product_id") REFERENCES "products"("id") ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "product_adventage"(
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "text_uz" TEXT NOT NULL,
    "text_ru" TEXT NOT NULL,
    "text_en" TEXT NOT NULL,
    "product_id" UUID NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_product FOREIGN KEY ("product_id") REFERENCES "products"("id") ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "chemical_property" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "product_id" UUID NOT NULL,
    "name_uz" VARCHAR NOT NULL,
    "name_ru" VARCHAR NOT NULL,
    "name_en" VARCHAR NOT NULL,   
    "unit" VARCHAR(50),
    "standard_min" FLOAT,
    "standard_max" FLOAT,
    "analysis_result" FLOAT NOT NULL,
    CONSTRAINT fk_product FOREIGN KEY ("product_id") REFERENCES "products"("id") ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "corrosion_impact" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "product_id" UUID NOT NULL,
    "material_uz" VARCHAR(255) NOT NULL,
    "material_ru" VARCHAR(255) NOT NULL,
    "material_en" VARCHAR(255) NOT NULL,
    "unit" VARCHAR(50) DEFAULT 'g/cm³',
    "max_limit" FLOAT NOT NULL,
    "analysis_result" FLOAT NOT NULL,
    CONSTRAINT fk_product FOREIGN KEY ("product_id") REFERENCES "products"("id") ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "product_files"(
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "file" BYTEA NOT NULL,
    "product_id" UUID NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_product FOREIGN KEY ("product_id") REFERENCES "products"("id") ON DELETE CASCADE
);