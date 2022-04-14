### About

An email and sms notification system

### Preflight

The application will check for Env variables before booting up. The .env-example is used as a reference, please keep this up to date.

### Start Local Instance

Start up a local instance of the application, using docker build + serve

From terminal run: `task run`

### Run Testing Suite

From terminal run: `task integration-test`

### Operating System Tools Needed

- Go 1.16
- [Docker & Docker Compose](https://docs.docker.com/get-docker/)
- [Taskfile](https://taskfile.dev/#/installation)

### Overview

Http = internal/transport/http\
Service = internal/comment/comment\
Repository = internal/db

```mermaid
flowchart LR
  /api/v1/endpoint-name <--> Http <--> Service <--> Repository <--> Postgres
  Service <--> Client

```

```mermaid
sequenceDiagram
    cmd/server/main->>+db: NewDatabase()
    cmd/server/main->>+db: MigrateDb()
    db->>+cmd/server/main: *db
    cmd/server/main->>+internal/service: NewService(*db)
    internal/service->>-cmd/server/main: *service
    cmd/server/main->>+internal/transport/http: NewHandler(*service)
    internal/transport/http->>-cmd/server/main: serve http endpoints
```
