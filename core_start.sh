#!/bin/sh

set -e

docker run -p 8081:8081 --env TELEGRAM_API_ID=API_ID --env TELEGRAM_API_HASH=API_HASH tdlight/tdlightbotapi
docker-compose up -d