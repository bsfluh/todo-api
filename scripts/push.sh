#!/usr/bin/env bash

set -euo pipefail

echo "Проверяю изменения..."
git status

echo "Добавляю все файлы..."
git add .

echo "Введи сообщение коммита (например, 'Добавил handler для tasks'):"
read commit_msg

git commit -m "$commit_msg"

echo "Пушу в main..."
git push origin main

echo "Готово! Проверь на GitHub."