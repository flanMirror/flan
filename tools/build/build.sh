#!/bin/sh

go build -trimpath -ldflags "-s -w" -tags=jsoniter -o build/ $PWD/cmd/misskey
go build -trimpath -o build/ $PWD/cmd/prairie
