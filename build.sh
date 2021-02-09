#!/usr/bin/env bash

go-bindata -pkg  asset templates/...
mv bindata.go asset/
go build -o dockeryard cmd/main.go
