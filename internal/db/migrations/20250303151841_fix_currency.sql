BEGIN TRANSACTION;

UPDATE cashes SET currency = 99 WHERE currency = 2;
UPDATE cashes SET currency = 2 WHERE currency = 1;
UPDATE cashes SET currency = 1 WHERE currency = 0;

UPDATE categories SET currency = 99 WHERE currency = 2;
UPDATE categories SET currency = 2 WHERE currency = 1;
UPDATE categories SET currency = 1 WHERE currency = 0;

UPDATE items SET currency = 99 WHERE currency = 2;
UPDATE items SET currency = 2 WHERE currency = 1;
UPDATE items SET currency = 1 WHERE currency = 0;

COMMIT;
