DB_TYPE     ?= mysql
DB_NAME     ?= learning-golang-db
DB_USER		?= mysql-user
DB_PASSWORD ?= password
DB_SOURCE   ?= $(DB_USER):$(DB_PASSWORD)@/$(DB_NAME)
SERVER_PORT ?= 3000
GOARCH      := amd64
GOOS        := linux
CGO_ENABLED := 0
TAG         := latest
SEED_DATA_SOURCE := $(PWD)/db/seed.json

install_dependencies:
	GO111MODULE=on go get github.com/amacneil/dbmate
	GO111MODULE=on go get github.com/matryer/moq

generate_mocks:
	moq -out pkg/cart/repository_mock.go pkg/cart Repository

generate_seed_data:
	GO111MODULE=on go run ./cmd/lggenseeddata \
		--json-destination=$(SEED_DATA_SOURCE) \
		--item-count=100 \
		--manufacturer-count=5

test:
	DB_SOURCE=$(DB_SOURCE) \
	DB_TYPE=$(DB_TYPE) \
	SERVER_PORT=$(SERVER_PORT) \
	JSON_SOURCE=$(SEED_DATA_SOURCE) \
	GO111MODULE=on go test -race -v ./...

run_local_db:
	(docker rm -f $(DB_NAME) || true) && docker run -d \
		--name $(DB_NAME) \
		-e MYSQL_ROOT_PASSWORD=$(DB_PASSWORD) \
		-e MYSQL_USER=$(DB_USER) \
		-e MYSQL_PASSWORD=$(DB_PASSWORD) \
		-e MYSQL_DATABASE=$(DB_NAME) \
		-p 3306:3306 \
		mysql:8.0.16

run_migrations:
	DATABASE_URL=mysql://$(DB_SOURCE) dbmate up

seed:
	GO111MODULE=on go run ./cmd/lgseed \
		--db-source=$(DB_SOURCE) \
		--db-type=$(DB_TYPE) \
		--json-source=$(SEED_DATA_SOURCE)

build_server:
	GO111MODULE=on go build -o dist/lgserver ./cmd/lgserver

run_server: build_server
	./dist/lgserver \
		--db-source=$(DB_SOURCE) \
		--db-type=$(DB_TYPE) \
		--server-port=$(SERVER_PORT)

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
