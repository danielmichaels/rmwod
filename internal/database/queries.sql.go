// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: queries.sql

package database

import (
	"context"
	"database/sql"
)

const insertVideo = `-- name: InsertVideo :exec
INSERT INTO videos (title, url, description, duration)
VALUES (?, ?, ?, ?)
`

type InsertVideoParams struct {
	Title       string         `json:"title"`
	Url         string         `json:"url"`
	Description sql.NullString `json:"description"`
	Duration    int64          `json:"duration"`
}

func (q *Queries) InsertVideo(ctx context.Context, arg InsertVideoParams) error {
	_, err := q.db.ExecContext(ctx, insertVideo,
		arg.Title,
		arg.Url,
		arg.Description,
		arg.Duration,
	)
	return err
}

const randomWod = `-- name: RandomWod :one
SELECT id, title, url, description, duration, tag_id
FROM videos
ORDER BY RANDOM() LIMIT 1
`

func (q *Queries) RandomWod(ctx context.Context) (Videos, error) {
	row := q.db.QueryRowContext(ctx, randomWod)
	var i Videos
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Url,
		&i.Description,
		&i.Duration,
		&i.TagID,
	)
	return i, err
}
