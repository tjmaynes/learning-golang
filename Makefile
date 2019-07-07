DB_TYPE     ?= sqlite3
DB_SOURCE     ?= $(PWD)/db/learning-golang.db
SERVER_PORT ?= 3000
GOARCH      := amd64
GOOS        := linux
CGO_ENABLED := 0
TAG         := latest
SEED_DATA_SOURCE := $(PWD)/db/seed.json

install_dependencies:
	GO111MODULE=on go get github.com/amacneil/dbmate
	GO111MODULE=on go get github.com/matryer/moq
	GO111MODULE=on CGO_ENABLED=1 go get github.com/mattn/go-sqlite3

generate_mocks:
	moq -out pkg/cart/repository_mock.go pkg/cart Repository

generate_seed_data:
	GO111MODULE=on go run ./cmd/lggenseeddata \
		--seed-data-destination=$(SEED_DATA_SOURCE) \
		--item-count=100 \
		--manufacturer-count=5

test:
	DB_TYPE=$(DB_TYPE) \
	DB_SOURCE=$(DB_SOURCE) \
	SERVER_PORT=$(SERVER_PORT) \
	SEED_DATA_SOURCE=$(SEED_DATA_SOURCE) \
	GO111MODULE=on go test -race -v ./...

run_migrations:
	DATABASE_URL=sqlite:///$(DB_SOURCE) dbmate up

seed_db:
	GO111MODULE=on go run ./cmd/lgseed \
		--db-type=$(DB_TYPE) \
		--db-source=$(DB_SOURCE) \
		--seed-data-source=$(SEED_DATA_SOURCE)

build_server:
	GO111MODULE=on go build -o dist/lgserver ./cmd/lgserver

run_server: build_server
	./dist/lgserver \
		--db-type=$(DB_TYPE) \
		--db-source=$(DB_SOURCE) \
		--server-port=$(SERVER_PORT)

build_image:
	docker build -t tjmaynes/learning-golang-server:$(TAG) .

run_image:
	docker run --rm \
	 --env DB_TYPE=$(DB_TYPE) \
	 --env DB_SOURCE=$(DB_SOURCE) \
	 --env SERVER_PORT=$(SERVER_PORT) \
	 --volume $(PWD):. \
	 --publish $(SERVER_PORT):$(SERVER_PORT) \
	 tjmaynes/learning-golang-server:$(TAG)

clean:
	rm -rf dist/ vendor/
