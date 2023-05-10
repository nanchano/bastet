-- name: CreateEvent :one
INSERT INTO events (
  name, description, category, location, publisher, lineup, start_ts, end_ts
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: GetEvent :one
SELECT * FROM events
WHERE id = $1 LIMIT 1;

-- name: UpdateEvent :one
UPDATE events
SET
  name = COALESCE(sqlc.narg('name'), name),
  description = COALESCE(sqlc.narg('description'), description),
  category = COALESCE(sqlc.narg('category'), category),
  location = COALESCE(sqlc.narg('location'), location),
  publisher = COALESCE(sqlc.narg('publisher'), publisher),
  lineup = COALESCE(sqlc.narg('lineup'), lineup),
  start_ts = COALESCE(sqlc.narg('start_ts'), start_ts),
  end_ts = COALESCE(sqlc.narg('end_ts'), end_ts)
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = $1;
