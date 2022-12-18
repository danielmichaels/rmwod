// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package database

import (
	"database/sql"
)

type Tags struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Videos struct {
	ID          int64          `json:"id"`
	Title       string         `json:"title"`
	Url         string         `json:"url"`
	Description sql.NullString `json:"description"`
	Duration    int64          `json:"duration"`
	TagID       sql.NullInt64  `json:"tag_id"`
}

type VideosTags struct {
	TagID   sql.NullInt64 `json:"tag_id"`
	VideoID sql.NullInt64 `json:"video_id"`
}