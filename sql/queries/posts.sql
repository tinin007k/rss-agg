-- name: CreatePosts :one
INSERT INTO posts (id, url,title,description,published_at,created_at, updated_at,feed_id)
VALUES ($1, $2, $3, $4, $5,$6,$7,$8)
RETURNING *;

-- name: GetPosts :many
SELECT * FROM posts JOIN feed_follows ON posts.feed_id = feed_follows.feed_id 
WHERE feed_follows.feed_id=$1;

