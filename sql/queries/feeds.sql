-- name: CreateFeed :one
INSERT into feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING *;

-- name: GetFeedsByUser :many
SELECT * FROM feeds
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT 25;

-- name: GetAllFeeds :many
SELECT * FROM feeds
ORDER BY created_at DESC
LIMIT $1;

-- name: GetFeedByUrl :one
SELECT * FROM feeds
WHERE url = $1;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: DeleteAllFeeds :exec
DELETE FROM feeds;

-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;
