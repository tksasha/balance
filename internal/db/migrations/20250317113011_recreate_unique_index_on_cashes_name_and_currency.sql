BEGIN TRANSACTION;

DROP INDEX IF EXISTS index_cashes_on_name_and_currency;

CREATE UNIQUE INDEX "unique_index_on_cashes_name_and_currency" ON "cashes"(
  "name",
  "currency"
) WHERE deleted_at IS NULL;

COMMIT;
