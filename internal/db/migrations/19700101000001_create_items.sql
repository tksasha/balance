CREATE TABLE IF NOT EXISTS "items" (
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
    CONSTRAINT "fk_rails_89fb86dc8b" FOREIGN KEY ("category_id") REFERENCES "categories" ("id")
)
