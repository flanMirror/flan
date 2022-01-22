.NOTPARALLEL: assets build start
SHELL = sh

all: build

.PHONY: build
build:
	go build -o build/ $$PWD/cmd/misskey

.PHONY: assets
assets:
	# TODO

