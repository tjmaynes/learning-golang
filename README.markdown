# Learning Golang

![https://godoc.org/github.com/tjmaynes/learning-golang](https://github.com/golang/gddo/blob/c782c79e0a3c3282dacdaaebeff9e6fd99cb2919/gddo-server/assets/status.svg)

> CRUD service with PostgreSQL database calls. Based on this [tutorial](https://itnext.io/building-restful-web-api-service-using-golang-chi-mysql-d85f427dee54), added some complexity to further understand Golang. Ended up using PostgreSQL and consolidating `repository` and `models` into `posts`.

## Requirements

- [golang](https://golang.org/)
- [docker](https://hub.docker.com/_/postgres)

## Usage

To install project dependencies, run the following command:
```bash
make install_dependencies
```

To run all tests, run the following command:
```bash
make test
```

To run the local database, run the following command:
```bash
make run_local_db
```

To run migrations, run the following command:
```bash
make run_migrations
```

To seed the database, run the following command:
```bash
make run_seed_job
```

To build the server, run the following command:
```bash
make build_server
```

To run the server, run the following command:
```bash
DB_SOURCE=<some-database-source> \
make run_server
```
