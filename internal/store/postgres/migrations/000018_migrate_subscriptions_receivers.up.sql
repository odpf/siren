INSERT INTO subscriptions_receivers(subscription_id, receiver_id, created_at, updated_at) (
SELECT 
    s.id as subscription_id,
    (nr->'id')::bigint as receiver_id,
    s.created_at as created_at,
    s.updated_at as updated_at
FROM subscriptions s, jsonb_array_elements(s.receiver) as nr
INNER JOIN receivers r ON r.id = (nr->'id')::bigint)
ON CONFLICT DO NOTHING;