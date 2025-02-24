CREATE TABLE IF NOT EXISTS "partners"(
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "logo" BYTEA NOT NULL,
    "name" VARCHAR NOT NULL,
    "flag" BYTEA NOT NULL,
    "partner_name" VARCHAR,
    "email" VARCHAR,
    "phone_number" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)