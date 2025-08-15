#!/bin/bash

set -e

APP_NAME="jimcostdev-utils"
MAIN_FILE="main.go"

echo "Compilando para Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -o ${APP_NAME}-windows-amd64.exe ${MAIN_FILE}

echo "Compilando para Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -o ${APP_NAME}-linux-amd64 ${MAIN_FILE}

echo "Compilando para macOS (amd64)..."
GOOS=darwin GOARCH=amd64 go build -o ${APP_NAME}-darwin-amd64 ${MAIN_FILE}

echo "¡Binarios generados!"
