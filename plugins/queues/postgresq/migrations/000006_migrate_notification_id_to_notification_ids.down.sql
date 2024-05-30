UPDATE message_queue SET notification_id = notification_ids->>0 WHERE notification_ids IS NOT NULL AND notification_id IS NULL;

ALTER TABLE
  message_queue
ADD COLUMN IF NOT EXISTS notification_id text;

CREATE INDEX IF NOT EXISTS message_queue_notification_id_idx ON message_queue (notification_id);