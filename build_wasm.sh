#!/bin/sh

GOOS=js GOARCH=wasm go build -o webroot/main.wasm src/*.go