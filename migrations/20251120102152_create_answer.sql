-- +goose Up
-- +goose StatementBegin
CREATE TABLE app.answers (
    id SERIAL PRIMARY KEY,
    question_id INT NOT NULL REFERENCES app.questions(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES app.users(id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

COMMENT ON COLUMN app.answers.id IS 'Уникальный идентификатор ответа';
COMMENT ON COLUMN app.answers.question_id IS 'Ссылка на вопрос';
COMMENT ON COLUMN app.answers.user_id IS 'Ссылка на пользователя';
COMMENT ON COLUMN app.answers.text IS 'Текст ответа';
COMMENT ON COLUMN app.answers.created_at IS 'Время создания ответа';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE app.answers;
-- +goose StatementEnd