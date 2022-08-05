#!/bin/bash

go mod tidy && go mod vendor

go build -o app

/app/app