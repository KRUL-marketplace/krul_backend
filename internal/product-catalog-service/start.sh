#!/bin/bash

# Проверка наличия необходимых утилит
command -v docker >/dev/null 2>&1 || { echo >&2 "Docker не установлен. Установите Docker и попробуйте снова."; exit 1; }
command -v docker-compose >/dev/null 2>&1 || { echo >&2 "Docker Compose не установлен. Установите Docker Compose и попробуйте снова."; exit 1; }
command -v make >/dev/null 2>&1 || { echo >&2 "Make не установлен. Установите Make и попробуйте снова."; exit 1; }

# Создание .env файла на основе .env.example
if [ ! -f .env ]; then
  if [ -f .env.example ]; then
    cp .env.example .env
    echo ".env файл создан на основе .env.example"
  else
    echo "Файл .env.example не найден. Создайте .env файл вручную."
    exit 1
  fi
fi

# Выполнение make команд до generate (включительно)
make vendor-proto
make install-golangci-lint
make lint
make install-deps
make get-deps
make generate

# Сборка docker контейнера
sudo docker compose up --build

echo "Сборка завершена!"
