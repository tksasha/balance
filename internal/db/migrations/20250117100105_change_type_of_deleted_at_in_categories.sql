PRAGMA foreign_keys = OFF;

CREATE TABLE "categories2" (
    "id" integer NOT NULL PRIMARY KEY,
    "name" varchar DEFAULT NULL,
    "income" boolean DEFAULT 0,
    "visible" boolean DEFAULT 1,
    "currency" integer DEFAULT 0,
    "supercategory" integer DEFAULT 1 NOT NULL,
    "deleted_at" datetime
);

INSERT INTO
    categories2
SELECT
    id,
    name,
    income,
    visible,
    currency,
    supercategory,
    deleted_at
FROM
    categories;

DROP TABLE categories;

ALTER TABLE categories2
RENAME TO categories;

CREATE UNIQUE INDEX "index_categories_on_name_and_currency" ON "categories" ("name", "currency");

PRAGMA foreign_keys = OFF;
