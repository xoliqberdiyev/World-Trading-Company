CREATE TABLE IF NOT EXISTS "product_specification"(
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" VARCHAR NOT NULL,
    "brands" TEXT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "product_features"(
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "text" TEXT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)