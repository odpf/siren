CREATE TABLE IF NOT EXISTS subscriptions_receivers (
    id bigserial PRIMARY KEY,
    subscription_id bigint REFERENCES subscriptions(id),
    receiver_id bigint REFERENCES receivers(id),
    labels jsonb,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    deleted_at timestamptz,
    UNIQUE (subscription_id, receiver_id)
);

CREATE UNIQUE INDEX IF NOT EXISTS subscriptions_receivers_idx_subscription_receiver ON subscriptions_receivers(subscription_id, receiver_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS subscriptions_receivers_idx_receiver_id ON subscriptions_receivers(receiver_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS subscriptions_receivers_idx_labels ON subscriptions_receivers USING GIN(labels jsonb_path_ops) WHERE deleted_at IS NULL;