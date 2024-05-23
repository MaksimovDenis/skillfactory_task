include .env
LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

migration-status:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN_MAKEFILE} status -v

migration-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN_MAKEFILE} up -v

migration-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN_MAKEFILE} down -v

build:
	go mod download && CGD_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/main.go

up:
	docker-compose up --build 

down:
	docker-compose down
