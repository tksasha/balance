BEGIN TRANSACTION;

UPDATE items
SET
    category_slug = categories.slug
FROM
    categories
WHERE
    categories.id = items.category_id;

CREATE TABLE IF NOT EXISTS "items2" (
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
    CONSTRAINT "fk_rails_89fb86dc8b" FOREIGN KEY ("category_id") REFERENCES "categories" ("id")
);

INSERT INTO
    items2
SELECT
    *
FROM
    items;

DROP TABLE items;

ALTER TABLE items2
RENAME TO items;

CREATE INDEX "index_balans_items_on_date_and_category_id" ON "items" ("date", "category_id");

CREATE INDEX "index_balans_items_on_date" ON "items" ("date");

COMMIT;
