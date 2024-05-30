DROP INDEX IF EXISTS notification_ids_idx;

ALTER TABLE
  message_queue
DROP COLUMN IF EXISTS notification_ids;