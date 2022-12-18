-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS videos
(
    id          INTEGER primary key autoincrement,
    title       text    NOT NULL,
    url         text    NOT NULL,
    description text,
    duration    INTEGER NOT NULL,
    tag_id      INTEGER,
    FOREIGN KEY (tag_id) REFERENCES tags (id)
);
CREATE TABLE IF NOT EXISTS tags
(
    id   INTEGER primary key autoincrement,
    name text NOT NULL
);
CREATE TABLE IF NOT EXISTS videos_tags
(
    tag_id   INTEGER,
    video_id INTEGER,
    FOREIGN KEY (tag_id) REFERENCES tags (id),
    FOREIGN KEY (video_id) REFERENCES videos (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS videos;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS videos_tags;
-- +goose StatementEnd
