-- +goose Up
CREATE TABLE feeds(
    id UUID PRIMARY KEY,
    name text not null,
    url TEXT not null unique,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    user_id UUID not null REFERENCES users(id) on delete cascade
);

-- +goose Down
DROP TABLE feeds;