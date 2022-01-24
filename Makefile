.NOTPARALLEL: all boil build assets init-db clean-db start-db stop-db import-db sqlboiler sqlboiler-test
SHELL = sh

PG_CTL = pg_ctl -D build/postgres/db -o "-k $$PWD/build/postgres/sock -p 3002"
PSQL = psql -h $$PWD/build/postgres/sock -p 3002 -U $$(whoami)

all: build
boil: init-db start-db import-db sqlboiler sqlboiler-test stop-db clean-db

.PHONY: build
build:
	go build -o build/ $$PWD/cmd/misskey

.PHONY: assets
assets:
	# TODO

.PHONY: init-db
init-db:
	mkdir -p build/postgres/db
	initdb build/postgres/db
	mkdir build/postgres/sock

.PHONY: clean-db
clean-db:
	rm -rf build/postgres

.PHONY: start-db
start-db:
	$(PG_CTL) start

.PHONY: stop-db
stop-db:
	$(PG_CTL) stop

.PHONY: import-db
import-db:
	$(PSQL) postgres -c "create user misskey;"
	$(PSQL) postgres -c "create database misskey;"
	$(PSQL) misskey < $$PWD/spec/misskey.sql

.PHONY: sqlboiler
sqlboiler:
	-go install github.com/volatiletech/sqlboiler/v4@latest
	-go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
	-env PSQL_HOST=$$PWD/build/postgres/sock $$(go env GOPATH)/bin/sqlboiler -c sqlboiler.toml psql

.PHONY: sqlboiler-test
sqlboiler-test:
	$(PSQL) postgres -c "alter user misskey with superuser;"
	-env PSQL_HOST=$$PWD/build/postgres/sock go test $$PWD/db/models
