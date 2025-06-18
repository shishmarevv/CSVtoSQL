@echo off
setlocal enabledelayedexpansion

set BIN_NAME=csvtosql
set BUILD_DIR=build

if not exist %BUILD_DIR% mkdir %BUILD_DIR%

echo Building for Windows 64-bit...
set GOOS=windows
set GOARCH=amd64
call go build -o %BUILD_DIR%\%BIN_NAME%_windows_amd64.exe ./cmd/

echo Building for Linux 64-bit...
set GOOS=linux
set GOARCH=amd64
call go build -o %BUILD_DIR%\%BIN_NAME%_linux_amd64 ./cmd/

echo Building for MacOS 64-bit (Intel)...
set GOOS=darwin
set GOARCH=amd64
call go build -o %BUILD_DIR%\%BIN_NAME%_darwin_amd64 ./cmd/

echo Building for MacOS ARM64 (Apple Silicon)...
set GOOS=darwin
set GOARCH=arm64
call go build -o %BUILD_DIR%\%BIN_NAME%_darwin_arm64 ./cmd/

echo Building for Linux ARM64...
set GOOS=linux
set GOARCH=arm64
call go build -o %BUILD_DIR%\%BIN_NAME%_linux_arm64 ./cmd/

echo.
echo Bin files built in: %BUILD_DIR%\
dir /b /o %BUILD_DIR%\
endlocal

