-- +goose Up
ALTER TABLE tasks ADD COLUMN IF NOT EXISTS deadline TIMESTAMPTZ;

-- +goose Down
ALTER TABLE tasks DROP COLUMN IF EXISTS deadline;