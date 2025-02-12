BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS "cashes" (
    "id" integer NOT NULL PRIMARY KEY,
    "sum" decimal(10, 2) DEFAULT NULL,
    "name" varchar DEFAULT NULL,
    "deleted_at" time DEFAULT NULL,
    "formula" varchar DEFAULT NULL,
    "currency" integer DEFAULT 0,
    "supercategory" integer DEFAULT 1 NOT NULL,
    "favorite" boolean DEFAULT 0
);

CREATE UNIQUE INDEX IF NOT EXISTS "index_cashes_on_name_and_currency" ON "cashes" ("name", "currency");

COMMIT;
