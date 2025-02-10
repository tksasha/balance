ALTER TABLE cashes
DROP COLUMN deleted_at;

ALTER TABLE cashes
ADD COLUMN deleted_at DATETIME;
