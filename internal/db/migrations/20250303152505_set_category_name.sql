BEGIN TRANSACTION;

UPDATE items
SET
    category_name = categories.name
FROM
    categories
WHERE
    items.category_id = categories.id;

COMMIT;
