#!/usr/bin/env bash

go build -o dockyard_macos cmd/main.go
GOOS=linux go build -o dockyard_linux cmd/main.go
GOOS=windows go build -o dockyard.exe cmd/main.go