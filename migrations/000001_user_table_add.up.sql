CREATE TABLE IF NOT EXISTS "users"(
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "username" VARCHAR UNIQUE NOT NULL,
    "first_name" VARCHAR,
    "last_name" VARCHAR,
    "password"  VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)