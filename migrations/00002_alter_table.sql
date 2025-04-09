-- +goose Up
ALTER TABLE users ADD COLUMN status TEXT DEFAULT 'active';

-- +goose Down
ALTER TABLE users DROP COLUMN status;