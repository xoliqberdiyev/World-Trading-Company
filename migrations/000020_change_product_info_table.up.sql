ALTER TABLE "chemical_property" ADD COLUMN "standard_range" VARCHAR(100);
ALTER TABLE "chemical_property" DROP COLUMN "standard_min";
ALTER TABLE "chemical_property" DROP COLUMN "standard_max";
ALTER TABLE "corrosion_impact"
ALTER COLUMN "max_limit" TYPE VARCHAR(100) USING "max_limit"::VARCHAR;
