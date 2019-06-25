DB_TYPE     ?= postgres
DB_NAME     ?= learning-golang-db
DB_SOURCE   ?= postgres://postgres:password@localhost/$(DB_NAME)?sslmode=disable
SERVER_PORT ?= 3000
GOARCH      := amd64
GOOS        := linux
CGO_ENABLED := 0
TAG         := latest
SEED_DATA_SOURCE := ./cmd/lgseed/data.json

install_dependencies:
	go get -u github.com/amacneil/dbmate

test:
	GO111MODULE=on go test -v ./...

run_local_db:
	(docker rm -f $(DB_NAME) || true) && docker run -d \
		--name $(DB_NAME) \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=password \
		-p 5432:5432 \
		postgres:9.5.14-alpine

run_migrations:
	DATABASE_URL=$(DB_SOURCE) dbmate up

run_seed:
	DB_SOURCE=$(DB_SOURCE) \
	DB_TYPE=$(DB_TYPE) \
	JSON_SOURCE=$(SEED_DATA_SOURCE) \
	GO111MODULE=on go run ./cmd/lgseed

build_server:
	GO111MODULE=on go build -o dist/lgserver ./cmd/lgserver

run_server: build_server
	DB_SOURCE=$(DB_SOURCE) \
	DB_TYPE=$(DB_TYPE) \
	SERVER_PORT=$(SERVER_PORT) \
	./dist/lgserver

build_image:
	docker build -t tjmaynes/learning-golang-server:$(TAG) .

run_image:
	docker run --rm \
	 -p 3000:3000 \
	 -e DB_SOURCE=$(DB_SOURCE) \
	 -e DB_TYPE=$(DB_TYPE) \
	 -e SERVER_PORT=$(SERVER_PORT) \
	 tjmaynes/learning-golang-server:$(TAG)

clean:
	rm -rf dist/ vendor/
