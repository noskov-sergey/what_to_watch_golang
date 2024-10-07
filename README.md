# what_to_watch_golang

Проект написан для практики в принципах SOLID, а также использования пакета http/templates.
Данный проект в свое время был написан на python с использованием библиотеки Flask, 
но в ради эксперимента переписан на Golang

### Как запустить проект:

Cоздать файл с виртуальными переменными:

```
PG_DATABASE_NAME=xxxx
PG_USER=xxxx
PG_PASSWORD=xxxx
PG_PORT=5432
PG_DSN="host=localhost port=54221 dbname=opinion user=opinion-user password=opinion-password"

MIGRATION_DIR=./migrations

DB_DSN="postgres://user:password@localhost:5432/base-name?sslmode=disable"
BASE_URL="localhost:8080"
```

Запустить docker compose:

```
docker-compose up
```

Установить goose - утилита для выполенния миграций

```
make install-deps
```

И выполнить миграции

```
make local-migration-up
```

Запусить сервер

```
go run main.go
```