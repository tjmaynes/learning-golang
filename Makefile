DB_SOURCE   ?= "postgres://postgres:password@localhost/learning-golang?sslmode=disable"
DB_TYPE     ?= "postgres"
DB_NAME     ?= "learning-golang-db"
SERVER_PORT ?= "3000"
GOARCH      := "amd64"
GOOS        := "linux"
CGO_ENABLED := 0
TAG         := latest

run_local_db:
	(docker rm -f $(DB_NAME) || true) && docker run -d \
		--name $(DB_NAME) \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=password \
		-p 5432:5432 \
		postgres:9.5.14-alpine

run_server: build_server
	DB_SOURCE=$(DB_SOURCE) \
	DB_TYPE=$(DB_TYPE) \
	SERVER_PORT=$(SERVER_PORT) \
	./dist/learning-golang

build_server:
	GO111MODULE=on go build -o dist/lgserver ./cmd/lgserver

build_image:
	docker build -t tjmaynes/learning-golang-server:$(TAG) .

run_image:
	docker run --rm \
	 -p 3000:3000 \
	 -e DB_SOURCE=$(DB_SOURCE) \
	 -e DB_TYPE=$(DB_TYPE) \
	 -e SERVER_PORT=$(SERVER_PORT) \
	 tjmaynes/learning-golang-server:$(TAG)

run_migrations:
	DATABASE_URL=$(DB_SOURCE) dbmate up

run_seed:
	DB_SOURCE=$(DB_SOURCE) \
	DB_TYPE=$(DB_TYPE) \
	SERVER_PORT=$(SERVER_PORT) \
	GO111MODULE=on go run ./cmd/lgseed

clean:
	rm -rf dist/ vendor/