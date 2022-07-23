#!/bin/bash

go mod tidy
go mod vendor

CompileDaemon --build="go build -o app" --command=./app
