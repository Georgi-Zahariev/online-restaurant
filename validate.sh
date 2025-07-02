#!/usr/bin/env bash
set -euo pipefail

# Allow overriding via env, else defaults:
HOST=${DB_HOST:-localhost}
PORT=${DB_PORT:-5432}
DB=${DB_NAME:-restaurant}
USER=${DB_USER:-postgres}
PASS=${DB_PASSWORD:-postgres}
SSL=${DB_SSLMODE:-disable}

DB_URL="postgres://${USER}:${PASS}@${HOST}:${PORT}/${DB}?sslmode=${SSL}"

echo "Reverting all migrations to version 0..."
migrate -path db/migrations -database "$DB_URL" goto 0

echo "Applying all migrations..."
migrate -path db/migrations -database "$DB_URL" up

echo "Migration validation complete."
