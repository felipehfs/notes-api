-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS users(
    id VARCHAR(64) NOT NULL PRIMARY KEY,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    active BOOLEAN
);

CREATE TABLE IF NOT EXISTS notes (
    id VARCHAR(64) NOT NULL PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    description TEXT NOT NULL,
    ownerId VARCHAR(64) NOT NULL,
    created_at DATETIME default current_timestamp,
    last_updated DATETIME default current_timestamp,
    FOREIGN KEY (ownerId) REFERENCES users(id)
    ON DELETE SET NULL
);


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE users;
DROP TABLE notes;