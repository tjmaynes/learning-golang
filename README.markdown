# Learning Golang

> CRUD service with PostgreSQL database calls. Based on this [tutorial](https://itnext.io/building-restful-web-api-service-using-golang-chi-mysql-d85f427dee54), added some complexity to further understand Golang. Ended up using PostgreSQL and consolidating `repository` and `models` into `posts`.

## Requirements

- [golang](https://golang.org/)
- [docker](https://hub.docker.com/_/postgres)

## Usage

To run the local database, run the following command:
```bash
make run_local_db
```

To run the server, run the following command:
```bash
DB_SOURCE=<some-database-source> \
make run_server
```

To build the server, run the following command:
```bash
make build_server
```

To run migrations, run the following command:
```bash
make run_migrations
```

## Notes

### Context API
