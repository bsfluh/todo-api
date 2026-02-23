-- +goose Up

CREATE TABLE IF NOT EXISTS tasks (
    id            BIGSERIAL PRIMARY KEY,
    user_id       BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title         TEXT NOT NULL CHECK (length(title) >= 1),
    priority      TEXT NOT NULL CHECK (priority IN ('Low', 'Middle', 'High', 'Paramount')),
    created_time  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    update_time   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_tasks_user_id_priority ON tasks(user_id, priority);

-- +goose Down
DROP TABLE IF EXISTS tasks;