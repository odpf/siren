UPDATE message_queue SET notification_ids = json_build_array(notification_id) WHERE notification_id IS NOT NULL AND notification_ids IS NULL;

DROP INDEX IF EXISTS message_queue_notification_id_idx;

ALTER TABLE
  message_queue
DROP COLUMN IF EXISTS notification_id;