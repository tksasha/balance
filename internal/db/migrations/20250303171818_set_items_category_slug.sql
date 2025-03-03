BEGIN TRANSACTION;

UPDATE items
SET
    category_slug = categories.slug
FROM
    categories
WHERE
    categories.id = items.category_id;

COMMIT;
