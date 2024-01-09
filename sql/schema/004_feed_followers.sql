-- +goose Up
CREATE TABLE feed_followers (
    id UUID NOT NULL PRIMARY KEY,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    UNIQUE (feed_id, user_id)
);

-- +goose Down
DROP TABLE feed_followers;