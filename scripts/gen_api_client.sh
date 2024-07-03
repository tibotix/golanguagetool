#!/usr/bin/env bash

if ! [ -x "$(command -v swagger)" ]; then
    echo "go-swagger is not installed."
    echo "Please install it using `go install github.com/go-swagger/go-swagger/cmd/swagger@latest`"
    exit 1
fi

cd "$(dirname "$0")/../pkg"
swagger generate client -c api -f "https://languagetool.org/http-api/languagetool-swagger.json"
go mod tidy