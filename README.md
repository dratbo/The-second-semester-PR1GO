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
