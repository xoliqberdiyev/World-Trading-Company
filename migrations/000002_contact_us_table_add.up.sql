CREATE TABLE IF NOT EXISTS "contact_us"(
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),    
    "first_name" VARCHAR(250) NOT NULL,
    "last_name" VARCHAR(250) NOT NULL,
    "email" VARCHAR NOT NULL,
    "is_contacted" BOOLEAN DEFAULT false,
    "comment" TEXT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)