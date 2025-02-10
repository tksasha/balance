CREATE TABLE IF NOT EXISTS "schema_migrations" ("version" varchar NOT NULL);

CREATE UNIQUE INDEX IF NOT EXISTS "unique_schema_migrations" ON "schema_migrations" ("version");
