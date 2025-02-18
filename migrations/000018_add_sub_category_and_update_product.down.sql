DROP TABLE IF EXISTS "sub_categories";

ALTER TABLE "products" 
DROP CONSTRAINT fk_sub_category,
DROP COLUMN "sub_category_id";