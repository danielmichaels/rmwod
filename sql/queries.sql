-- name: RandomWod :one
SELECT *
FROM videos
ORDER BY RANDOM() LIMIT 1;

-- name: InsertVideo :exec
INSERT INTO videos (title, url, description, duration)
VALUES (?, ?, ?, ?);
