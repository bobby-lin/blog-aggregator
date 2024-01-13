-- +goose Up
CREATE TABLE posts (
    id UUID NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title VARCHAR(1024) NOT NULL,
    url VARCHAR(1024) NOT NULL UNIQUE,
    description VARCHAR(2048) NOT NULL,
    published_at TIMESTAMP,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;