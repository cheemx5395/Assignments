-- +goose Up
CREATE TABLE users(
    email TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;