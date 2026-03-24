-- +goose Up
ALTER TABLE tasks 
ADD COLUMN IF NOT EXISTS description TEXT;

-- +goose Down
ALTER TABLE tasks 
DROP COLUMN IF EXISTS description;