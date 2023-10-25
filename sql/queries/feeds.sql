-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, created_at, updated_at,user_id)
VALUES ($1, $2, $3, $4, $5,$6)
RETURNING *;

-- name: GetFeed :many
SELECT * FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT $1;


-- name: UpdateLastFetched :one
UPDATE feeds SET last_fetched_at=NOW(),updated_at=NOW() where id=$1 RETURNING *;