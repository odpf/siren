ALTER TABLE
  message_queue
ADD COLUMN IF NOT EXISTS notification_ids jsonb;

CREATE INDEX IF NOT EXISTS notification_ids_idx ON message_queue USING gin (notification_ids);