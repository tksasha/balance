CREATE TABLE IF NOT EXISTS "categories" (
    "id" integer NOT NULL PRIMARY KEY,
    "name" varchar DEFAULT NULL,
    "income" boolean DEFAULT 0,
    "visible" boolean DEFAULT 1,
    "currency" integer DEFAULT 0,
    "supercategory" integer DEFAULT 1 NOT NULL,
    "deleted_at" datetime (6)
)
