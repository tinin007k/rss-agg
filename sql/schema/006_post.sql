-- +goose Up
CREATE TABLE posts(
    id UUID PRIMARY KEY,
    url TEXT not null unique,
    title TEXT NOT NULL,
    description TEXT,
    published_at TIMESTAMP,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    feed_id UUID not null REFERENCES feeds(id) on delete cascade
);

-- +goose Down
DROP TABLE posts;