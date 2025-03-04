CREATE TABLE sqlite_sequence(name,seq);
CREATE TABLE IF NOT EXISTS "schema_migrations"("version" varchar NOT NULL);
CREATE UNIQUE INDEX "unique_schema_migrations" ON "schema_migrations"(
  "version"
);
CREATE TABLE IF NOT EXISTS "ar_internal_metadata"(
  "key" varchar NOT NULL PRIMARY KEY,
  "value" varchar,
  "created_at" datetime NOT NULL,
  "updated_at" datetime NOT NULL
);
CREATE TABLE IF NOT EXISTS "versions"(
  "id" integer PRIMARY KEY AUTOINCREMENT NOT NULL,
  "item_type" varchar NOT NULL,
  "item_id" integer(8) NOT NULL,
  "event" varchar NOT NULL,
  "whodunnit" varchar,
  "object" text(1073741823),
  "created_at" datetime,
  "object_changes" text(1073741823)
);
CREATE INDEX "index_versions_on_item_type_and_item_id" ON "versions"(
  "item_type",
  "item_id"
);
CREATE TABLE IF NOT EXISTS "cashes"(
  "id" integer NOT NULL PRIMARY KEY,
  "sum" decimal(10,2) DEFAULT NULL,
  "name" varchar DEFAULT NULL,
  "formula" varchar DEFAULT NULL,
  "currency" integer DEFAULT 0,
  "supercategory" integer DEFAULT 1 NOT NULL,
  "favorite" boolean DEFAULT 0,
  deleted_at DATETIME
);
CREATE UNIQUE INDEX "index_cashes_on_name_and_currency" ON "cashes"(
  "name",
  "currency"
);
CREATE TABLE IF NOT EXISTS "categories"(
  "id" integer NOT NULL PRIMARY KEY,
  "name" varchar DEFAULT NULL,
  "income" boolean DEFAULT 0,
  "visible" boolean DEFAULT 1,
  "currency" integer DEFAULT 0,
  "supercategory" integer DEFAULT 1 NOT NULL,
  "deleted_at" datetime
  ,
  slug VARCHAR
);
CREATE UNIQUE INDEX "index_categories_on_name_and_currency" ON "categories"(
  "name",
  "currency"
);
CREATE TABLE IF NOT EXISTS "items"(
  "id" integer NOT NULL PRIMARY KEY,
  "date" date DEFAULT NULL,
  "sum" decimal(10, 2) NOT NULL,
  "category_id" integer DEFAULT NULL,
  "description" varchar DEFAULT NULL,
  "created_at" datetime DEFAULT NULL,
  "updated_at" datetime DEFAULT NULL,
  "formula" text DEFAULT NULL,
  "deleted_at" time DEFAULT NULL,
  "currency" integer DEFAULT 0,
  category_name VARCHAR NOT NULL,
  category_slug VARCHAR NOT NULL,
  CONSTRAINT "fk_rails_89fb86dc8b" FOREIGN KEY("category_id") REFERENCES "categories"("id")
);
CREATE INDEX "index_balans_items_on_date_and_category_id" ON "items"(
  "date",
  "category_id"
);
CREATE INDEX "index_balans_items_on_date" ON "items"("date");
