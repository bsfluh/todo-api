#!/usr/bin/env bash

set -euo pipefail

echo "Подтягиваю изменения из main..."
git pull origin main

echo "Проверь, что появилось:"
ls -la

echo "Готово!"