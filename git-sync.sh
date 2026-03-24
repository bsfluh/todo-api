#!/usr/bin/env bash
# ================================================
# Короткие команды: push, pull, sync
# ================================================

cd /c/Users/pgar7/todo-api || { echo "❌ Папка не найдена!"; exit 1; }

# Авто ssh-agent + ключ
if [ -z "$SSH_AUTH_SOCK" ]; then
    eval "$(ssh-agent -s)" >/dev/null 2>&1
fi
ssh-add ~/.ssh/id_ed25519_bsfluh_win >/dev/null 2>&1 || true

case "$1" in
  pull)
    echo "📥 pull..."
    git pull origin main
    ;;

  push)
    echo "📤 push..."
    if ! git diff --quiet || ! git diff --cached --quiet || [ -n "$(git ls-files --others --exclude-standard)" ]; then
        git add .
        git commit -m "Auto commit: $(date '+%Y-%m-%d %H:%M:%S')"
        git push origin main
    else
        echo "✅ Нет изменений для пуша"
    fi
    ;;

  sync)
    echo "🔄 sync (pull + push)..."
    git pull origin main
    if ! git diff --quiet || ! git diff --cached --quiet || [ -n "$(git ls-files --others --exclude-standard)" ]; then
        git add .
        git commit -m "Auto sync: $(date '+%Y-%m-%d %H:%M:%S')"
        git push origin main
    else
        echo "✅ Нет новых изменений"
    fi
    ;;

  *)
    echo "Использование:"
    echo "   push     → закоммитить и отправить"
    echo "   pull     → скачать изменения"
    echo "   sync     → pull + push"
    ;;
esac