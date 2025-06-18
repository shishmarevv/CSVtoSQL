#!/bin/bash
set -e

echo "CSVtoSQL bin file build started"

BIN_NAME=csvtosql
BUILD_DIR=build
mkdir -p $BUILD_DIR

# Windows 64-bit
GOOS=windows GOARCH=amd64 go build -o $BUILD_DIR/${BIN_NAME}_windows_amd64.exe ./cmd/
# Linux 64-bit
GOOS=linux GOARCH=amd64 go build -o $BUILD_DIR/${BIN_NAME}_linux_amd64 ./cmd/
# MacOS 64-bit (Intel)
GOOS=darwin GOARCH=amd64 go build -o $BUILD_DIR/${BIN_NAME}_darwin_amd64 ./cmd/


# MacOS ARM64 (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o $BUILD_DIR/${BIN_NAME}_darwin_arm64 ./cmd/
# Linux ARM64
GOOS=linux GOARCH=arm64 go build -o $BUILD_DIR/${BIN_NAME}_linux_arm64 ./cmd/

echo "Bin files built in: $BUILD_DIR:"
ls -lh $BUILD_DIR