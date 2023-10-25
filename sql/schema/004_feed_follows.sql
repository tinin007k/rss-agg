-- +goose Up
CREATE TABLE feed_follows(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    user_id UUID not null REFERENCES users(id) on delete cascade,
    feed_id UUID not null REFERENCES feeds(id) on delete cascade,
    UNIQUE(user_id,feed_id)
);

-- +goose Down
DROP TABLE feed_follows;