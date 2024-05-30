#!/usr/bin/bash

if [ -z "$1" ]; then
    exit 1
fi

go run cmd/$1/main.go
