#!/usr/bin/env bash
air
set -euo pipefail #для безопасности, чтобы скрипт не продолжал выполняться при ошибке, неиспользуемой переменной или ошибке в пайпе
echo "Starting server in development mode..."
go run backend/cmd/main.go