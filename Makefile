# not sure if this is still necessary but accidentally running them parallel can definitely cause a huge mess
.NOTPARALLEL: all full boil assets build static assets-package template init-db clean-db start-db stop-db import-db sqlboiler sqlboiler-test
SHELL = sh

TARGET = $(shell ./build/misskey target)
PG_CTL = pg_ctl -D build/postgres/db -o "-k $$PWD/build/postgres/sock -p 3002"
PSQL = psql -h $$PWD/build/postgres/sock -p 3002 -U $$(whoami)
ASSETS = cmd/misskey/assets
STATIC = $(ASSETS)/static

all: build
full: boil build
boil: init-db start-db import-db sqlboiler sqlboiler-test stop-db clean-db
assets: static template assets-package

.PHONY: build
build:
	go build -o build/ $$PWD/cmd/misskey
	go build -o build/ $$PWD/cmd/prairie

.PHONY: static
static:
	rm -rf $(STATIC)
	mkdir $(STATIC)

	# general files
	cp -r $$MISSKEY_ROOT/packages/backend/assets $(STATIC)/static-assets
	cp -r $$MISSKEY_ROOT/packages/client/assets $(STATIC)/client-assets
	cp -r $$MISSKEY_ROOT/built/_client_dist_ $(STATIC)/assets

	# special cases
	cp $$MISSKEY_ROOT/packages/backend/assets/apple-touch-icon.png $(STATIC)
	cp $$MISSKEY_ROOT/packages/backend/assets/robots.txt $(STATIC)
	cp $$MISSKEY_ROOT/built/_client_dist_/sw.$(TARGET).js $(STATIC)/sw.js
	cp -r $$MISSKEY_ROOT/packages/client/node_modules/@discordapp/twemoji/dist/svg $(STATIC)/twemoji
	# manifest TODO

.PHONY: assets-package
assets-package:
	# TODO

.PHONY: template
template:
	./build/prairie -w -o $(ASSETS)/template -i $$MISSKEY_ROOT

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
