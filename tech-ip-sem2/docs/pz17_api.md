# API Documentation for Practical Work #1 (pz17)

## 1. Общая информация

Проект состоит из двух микросервисов:

- **Auth service** — отвечает за выдачу и проверку токенов (учебная упрощённая версия).
- **Tasks service** — управляет задачами (CRUD) с проверкой токена через Auth.

### Базовые адреса (по умолчанию)

| Сервис | Адрес |
|--------|-------|
| Auth   | `http://localhost:8081` |
| Tasks  | `http://localhost:8082` |

### Переменные окружения для запуска

**Auth service**
- `AUTH_PORT` — порт, на котором слушает Auth (по умолчанию `8081`).

**Tasks service**
- `TASKS_PORT` — порт Tasks (по умолчанию `8082`).
- `AUTH_BASE_URL` — базовый URL для доступа к Auth (по умолчанию `http://localhost:8081`).

---

## 2. Auth Service

### 2.1. Получение токена

**Endpoint:** `POST /v1/auth/login`

**Описание:** Выдаёт фиксированный токен `demo-token` при любых непустых логине и пароле (учебный пример).

**Заголовки:**
- `Content-Type: application/json` (обязательно)

**Тело запроса (JSON):**
```json
{
  "username": "student",
  "password": "student"
}
```
Успешный ответ (200 OK):

```json
{
"access_token": "demo-token",
"token_type": "Bearer"
}
```
Возможные ошибки:

| Код | Описание                                | Пример тела ответа  |
|-----|-----------------------------------------|---------------------|
| 400 | `Неверный формат JSON или пустые поля	` | `Invalid JSON`      |
| 405 | `Метод не поддерживается	`              | `Method not allowed`|
| 500 | `Внутренняя ошибка сервера	`            | `Could not generate token`|

### 2.1. Проверка токена
Endpoint: GET /v1/auth/verify

Описание: Проверяет валидность переданного токена. В учебной реализации единственный валидный токен — demo-token.

Заголовки:

- Authorization: Bearer <token> (обязательно)

- X-Request-ID (опционально) — для сквозной трассировки.

Успешный ответ (200 OK):
```json
{
"valid": true,
"subject": "student"
}
```
Ответ при невалидном токене (401 Unauthorized):
```json
{
"valid": false,
"error": "unauthorized"
}
```
Возможные ошибки:

| Код | Описание                     | Пример тела ответа  |
|-----|------------------------------|---------------------|
| 401 | `Отсутствует заголовок Authorization` | `{"valid":false,"error":"missing authorization header"}`      |
| 401 | `Неверный формат (не Bearer)`   | `{"valid":false,"error":"invalid authorization format"}`|
| 401 | `Невалидный токен`        | `{"valid":false,"error":"invalid token"}`|
| 405 | `Метод не поддерживается` | `"Method not allowed"`|

## 3. Tasks Service

Все эндпоинты Tasks требуют наличия заголовка `Authorization: Bearer <token>`.
Если токен отсутствует или невалиден, сервис возвращает `401 Unauthorized`. 
При недоступности Auth (или ответе 5xx) возвращается `503 Service Unavailable`.

### 3.1. Создание задачи

Endpoint: `POST /v1/tasks`

Описание: Создаёт новую задачу. ID генерируется автоматически, поля `created_at` и `updated_at` заполняются сервером.

Заголовки:

- Authorization: Bearer <token> (обязательно)

- Content-Type: application/json (обязательно)

- X-Request-ID (опционально)

Тело запроса (JSON):

```json
{
"title": "Начать дулать вторую практику без гпт",
"description": "шутка",
"due_date": "2026-01-10"
}
```
Поле `due_date` ожидается в формате `YYYY-MM-DD`. Поля `description` и `due_date` необязательны.

Успешный ответ (201 Created):

```json
{
"id": "t_001",
"title": "Начать дулать вторую практику без гпт",
"description": "шутка",
"due_date": "2026-01-10",
"done": false,
"created_at": "2026-03-06T12:00:00Z",
"updated_at": "2026-03-06T12:00:00Z"
}
```
Возможные ошибки:


| Код | Описание                                                |
|-----|---------------------------------------------------------|
| 400 | `Отсутствует обязательное поле title или неверный JSON` |
| 401 | `Отсутствует или невалидный токен`                           | 
| 503 | `Auth service недоступен`                                      | 
| 500 | `Внутренняя ошибка сервера`                               |

### 3.2. Получение списка задач

Endpoint: `GET /v1/tasks`

Описание: Возвращает все сохранённые задачи.

Заголовки:

- Authorization: Bearer <token> (обязательно)

- X-Request-ID (опционально)

Успешный ответ (200 OK):

```json
[
{
"id": "t_001",
"title": "Начать дулать вторую практику без гпт",
"description": "шутка",
"due_date": "2026-01-10",
"done": false,
"created_at": "2026-03-06T12:00:00Z",
"updated_at": "2026-03-06T12:00:00Z"
}
]
```
Возможные ошибки: те же, что и для создания (401, 503, 500).

### 3.3. Получение задачи по ID

Endpoint: `GET /v1/tasks/{id}`

Описание: Возвращает задачу с указанным идентификатором.

Параметры пути: `id` (строка, например `t_001`)

Заголовки:

- Authorization: Bearer <token> (обязательно)

- X-Request-ID (опционально)

Успешный ответ (200 OK):

```json
{
"id": "t_001",
"title": "Начать дулать вторую практику без гпт",
"description": "шутка",
"due_date": "2026-01-10",
"done": false,
"created_at": "2026-03-06T12:00:00Z",
"updated_at": "2026-03-06T12:00:00Z"
}
```

Устал писать возможные ошибки, сами разберётесь!

### 3.4. Обновление задачи

Endpoint: `PATCH /v1/tasks/{id}`

Описание: Частично обновляет поля существующей задачи. Передавать можно только те поля, которые нужно изменить.

Параметры пути: `id`

Заголовки:

- Authorization: Bearer <token> (обязательно)

- Content-Type: application/json (обязательно)

- X-Request-ID (опционально)

Тело запроса (JSON):

```json
{
"title": "Updated title",
"done": true
}
```
Можно передавать любые из полей: `title`, `description`, `due_date`, `done`. Отсутствующие поля не изменяются.

Успешный ответ (200 OK):

```json
{
"id": "t_001",
"title": "Начать дулать вторую практику без гпт",
"description": "шутка",
"due_date": "2026-01-10",
"done": true,
"created_at": "2026-03-06T12:00:00Z",
"updated_at": "2026-03-06T13:30:00Z"
}
```
### 3.5. Удаление задачи
Endpoint: `DELETE /v1/tasks/{id}`

Описание: Удаляет задачу по идентификатору.

Параметры пути: `id`

Заголовки:

- Authorization: Bearer <token> (обязательно)

- X-Request-ID (опционально)

Успешный ответ (204 No Content) — тело ответа отсутствует.

## 4. Проброс request-id (сквозная трассировка)

Все сервисы поддерживают заголовок `X-Request-ID`. 
Tasks пробрасывает этот заголовок при вызове Auth, что позволяет связать логи обоих сервисов.

Пример запроса с request-id:

```bash
curl -X POST http://localhost:8082/v1/tasks \
-H "Authorization: Bearer demo-token" \
-H "Content-Type: application/json" \
-H "X-Request-ID: my-unique-req-id" \
-d '{"title":"Test"}'
```
Пример логов (после выполнения запроса):

Tasks service:

```text
[my-unique-req-id] POST /v1/tasks 201 15ms
[my-unique-req-id] GET /v1/auth/verify 200 3ms
```
Auth service:

```text
[my-unique-req-id] GET /v1/auth/verify 200 1ms
```

## 5. Инструкция по запуску (Windows, PowerShell)
Предварительные требования
Установленный Go (версия 1.22+)

Git (опционально) (я не знаю почему гпт везде пишет слово `опционально`)

Postman или curl для тестирования

### Шаги
Клонировать репозиторий (если ещё не сделано):

```bash
git clone <url-репозитория>
cd tech-ip-sem2
```
Инициализировать модуль и установить зависимости:

```powershell
go mod init tech-ip-sem2
go get github.com/google/uuid
go mod tidy
```
Запустить Auth сервис (в отдельном окне PowerShell):

```powershell
cd services/auth
$env:AUTH_PORT = "8081"   # можно не указывать, по умолчанию 8081
go run ./cmd/auth
```
Запустить Tasks сервис (в другом окне PowerShell):

```powershell
cd services/tasks
$env:TASKS_PORT = "8082"
$env:AUTH_BASE_URL = "http://localhost:8081"
go run ./cmd/tasks
```
Проверить работу (через Postman или curl), используя эндпоинты, описанные выше.

Остановка сервисов — нажать Ctrl+C в каждом окне.

всё,устал копировать:)

