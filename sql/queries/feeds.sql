-- name: CreateFeed :one
INSERT into feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING *;

-- name: GetFeedsByUser :many
SELECT * FROM feeds
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT 25;

-- name: DeleteAllFeeds :exec
DELETE FROM feeds;
