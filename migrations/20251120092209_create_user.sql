-- +goose Up
-- +goose StatementBegin
CREATE TABLE app.users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    login TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);

COMMENT ON COLUMN app.users.id IS 'Уникальный идентификатор пользователя';
COMMENT ON COLUMN app.users.login IS 'Логин пользователя';
COMMENT ON COLUMN app.users.password IS 'Хэш пароля пользователя';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE app.users;
-- +goose StatementEnd
