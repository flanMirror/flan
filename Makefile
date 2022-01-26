# not sure if this is still necessary but accidentally running them parallel can definitely cause a huge mess
.NOTPARALLEL: all full boil assets build static assets-package assets-download template init-db clean-db start-db stop-db import-db sqlboiler sqlboiler-test
SHELL = sh

TARGET != ./build/misskey target
PG_CTL = pg_ctl -D build/postgres/db -o "-k $$PWD/build/postgres/sock -p 3002"
PSQL = psql -h $$PWD/build/postgres/sock -p 3002 -U $$(whoami)
ASSETS = cmd/misskey/assets
PUBLIC = $(ASSETS)/public
ASSETS_URL = https://cronut.cafe/~rand/misskey/assets-$(TARGET).tar.gz

all: build
full: boil build
boil: init-db start-db import-db sqlboiler sqlboiler-test stop-db clean-db
assets: static template assets-package

.PHONY:

build: .PHONY
	go build -o build/ $$PWD/cmd/misskey
	go build -o build/ $$PWD/cmd/prairie

static: .PHONY
	rm -rf $(PUBLIC)
	mkdir $(PUBLIC)

	# general files
	cp -r $$MISSKEY_ROOT/packages/backend/assets $(PUBLIC)/static-assets
	cp -r $$MISSKEY_ROOT/packages/client/assets $(PUBLIC)/client-assets
	cp -r $$MISSKEY_ROOT/built/_client_dist_ $(PUBLIC)/assets

	# special cases
	cp $$MISSKEY_ROOT/packages/backend/assets/apple-touch-icon.png $(PUBLIC)
	cp $$MISSKEY_ROOT/packages/backend/assets/robots.txt $(PUBLIC)
	cp $$MISSKEY_ROOT/built/_client_dist_/sw.$(TARGET).js $(PUBLIC)/sw.js
	cp -r $$MISSKEY_ROOT/packages/client/node_modules/@discordapp/twemoji/dist/svg $(PUBLIC)/twemoji
	# manifest TODO

assets-package: .PHONY
	tar -zcvf build/assets-$(TARGET).tar.gz -C cmd/misskey/assets/ .

assets-download: .PHONY
	rm -rf $(ASSETS)/public $(ASSETS)/template
	if type fetch >/dev/null 2>/dev/null; then fetch -o build/assets $(ASSETS_URL); \
elif type curl >/dev/null 2>/dev/null; then curl $(ASSETS_URL) > build/assets; \
elif type wget >/dev/null 2>/dev/null; then wget -o build/assets $(ASSETS_URL) ; fi
	tar -zxvf build/assets -C cmd/misskey/assets
	rm build/assets

template: .PHONY
	./build/prairie -w -o $(ASSETS)/template -i $$MISSKEY_ROOT

init-db: .PHONY
	mkdir -p build/postgres/db
	initdb build/postgres/db
	mkdir build/postgres/sock

clean-db: .PHONY
	rm -rf build/postgres

start-db: .PHONY
	$(PG_CTL) start

stop-db: .PHONY
	$(PG_CTL) stop

import-db: .PHONY
	$(PSQL) postgres -c "create user misskey;"
	$(PSQL) postgres -c "create database misskey;"
	$(PSQL) misskey < $$PWD/spec/misskey.sql

sqlboiler: .PHONY
	-go install github.com/volatiletech/sqlboiler/v4@latest
	-go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
	-env PSQL_HOST=$$PWD/build/postgres/sock $$(go env GOPATH)/bin/sqlboiler -c sqlboiler.toml psql

sqlboiler-test: .PHONY
	$(PSQL) postgres -c "alter user misskey with superuser;"
	-env PSQL_HOST=$$PWD/build/postgres/sock go test $$PWD/db/models
