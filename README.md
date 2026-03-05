<h1 align="center"> Привет! Я <a target="_blank"> Кармеев Артур из группы ЭФМО-01-25 </a> 
<img src="https://github.com/blackcater/blackcater/raw/main/images/Hi.gif" height="32"/></h1>
<h3 align="center"> Данная практика была выполнена с божьей помощью! :dizzy_face: </h3>

Структура проекта:

    └tech-ip-sem2/
    ├── go.mod
    ├── go.sum
    ├── README.md
    ├── tech-ip-sem2/
    │   ├── go.mod
    │   ├── go.sum
    │   ├── shared/
    │   │   ├── middleware/
    │   │   │   ├── logging.go
    │   │   │   └── requestid.go
    │   │   └── httpx/
    │   │       └── client.go
    │   └── services/
    │       ├── tasks/
    │       │   ├── internal/
    │       │   │   ├── service/
    │       │   │   │   └── storage.go
    │       │   │   ├── http/
    │       │   │   │   ├── server.go
    │       │   │   │   └── handlers/
    │       │   │   │       └── tasks.go
    │       │   │   └── client/
    │       │   │       └── authclient/
    │       │   │           └── client.go
    │       │   └── cmd/
    │       │       └── tasks/
    │       │           └── main.go
    │       └── auth/
    │           ├── internal/
    │           │   ├── service/
    │           │   │   └── auth.go
    │           │   └── http/
    │           │       ├── server.go
    │           │       └── handlers/
    │           │           └── auth.go
    │           └── cmd/
    │               └── auth/
    │                   └── main.go
    ├── docs/
    │   └── pz17_api.md
    ├── .idea/
    │   ├── .gitignore
    │   ├── modules.xml
    │   ├── tech-ip-sem2.iml
    │   ├── vcs.xml
    │   └── workspace.xml
    ├── shared/
    │   ├── middleware/
    │   │   ├── logging.go
    │   │   └── requestid.go
    │   └── httpx/
    │       └── client.go
    └── services/
        ├── tasks/
        │   ├── internal/
        │   │   ├── service/
        │   │   │   └── storage.go
        │   │   ├── http/
        │   │   │   ├── server.go
        │   │   │   └── handlers/
        │   │   │       └── tasks.go
        │   │   └── client/
        │   │       └── authclient/
        │   │           └── client.go
        │   └── cmd/
        │       └── tasks/
        │           └── main.go
        └── auth/
            ├── internal/
            │   ├── service/
            │   │   └── auth.go
            │   └── http/
            │       ├── server.go
            │       └── handlers/
            │           └── auth.go
            └── cmd/
                └── auth/
                    └── main.go


## 1. Границы ответственности сервисов
- ```Auth service``` — отвечает за выдачу и проверку токенов.

```POST /v1/auth/login``` – принимает логин/пароль, возвращает фиксированный токен ```demo-token``` (учебное упрощение).

```GET /v1/auth/verify``` – проверяет токен из заголовка ```Authorization``` и возвращает статус валидности и имя субъекта (```student```).

- ```Tasks service``` — управляет задачами (```CRUD```).

Хранит задачи в памяти (in-memory).

Перед выполнением любой операции (кроме специально не защищённых) вызывает Auth для проверки токена.

При невалидном токене возвращает ```401 Unauthorized```, при недоступности Auth – ```503 Service Unavailable```.

## 2. Схема взаимодействия

```sequenceDiagram
    participant C as Client
    participant T as Tasks Service
    participant A as Auth Service

    C->>T: Запрос с Authorization: Bearer <token> + X-Request-ID
    T->>A: GET /v1/auth/verify (таймаут 3с, проброс X-Request-ID)
    A-->>T: 200 OK (valid) / 401 Unauthorized
    T-->>C: Результат операции (200/201/404/401/503...)
```
