CREATE TABLE IF NOT EXISTS "contact_us_footer"(
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "full_name" VARCHAR,
    "phone_number" VARCHAR,
    "email" VARCHAR,
    "is_contacted" BOOLEAN DEFAULT false,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
