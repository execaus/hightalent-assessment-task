# API-сервис для вопросов и ответов

## Функциональные требования

### Модели

**Question** – вопрос:

- `id`: int
- `text`: str (текст вопроса)
- `created_at`: datetime

**Answer** – ответ на вопрос:

- `id`: int
- `question_id`: int (ссылка на `Question`)
- `user_id`: str (идентификатор пользователя, например uuid)
- `text`: str (текст ответа)
- `created_at`: datetime

### Методы API

#### Вопросы (Questions)

- `GET /questions/` — список всех вопросов
- `POST /questions/` — создать новый вопрос
- `GET /questions/{id}` — получить вопрос и все ответы на него
- `DELETE /questions/{id}` — удалить вопрос (вместе с ответами)

#### Ответы (Answers)

- `POST /questions/{id}/answers/` — добавить ответ к вопросу
- `GET /answers/{id}` — получить конкретный ответ
- `DELETE /answers/{id}` — удалить ответ

### Логика

- Нельзя создать ответ к несуществующему вопросу.
- Один и тот же пользователь может оставлять несколько ответов на один вопрос.
- При удалении вопроса должны удаляться все его ответы (каскадно).

### Запуск проекта

1. Создайте файл конфигурации `config/config.yaml` со следующим содержимым:

```yaml
server:
  port: "8080"

database:
  host: "postgres"
  port: 5432
  user: "postgres"
  password: "1234"
  name: "app"
```

> Важно: `host` должен соответствовать имени сервиса PostgreSQL в `docker-compose.yaml`.

2. Определите переменные окружения в файле `.env`:

```env
JWT_SECRET_KEY=1234
DATABASE_PASSWORD=1234
```

3. Запустите проект через Docker Compose:

```bash
docker compose up --build
```

Сервер будет доступен на `http://localhost:8081`.