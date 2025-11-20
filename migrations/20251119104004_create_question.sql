-- +goose Up
-- +goose StatementBegin
CREATE TABLE app.questions (
    id SERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON COLUMN app.questions.id IS 'Уникальный идентификатор вопроса';
COMMENT ON COLUMN app.questions.text IS 'Текст вопроса';
COMMENT ON COLUMN app.questions.created_at IS 'Время создания вопроса';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE questions;
-- +goose StatementEnd
