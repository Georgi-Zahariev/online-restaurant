#!/usr/bin/env bash
set -euo pipefail

DB_URL="postgres://$USER@localhost:5432/restaurant?sslmode=disable"

echo "Rolling down previous migrations..."
migrate -path db/migrations -database "$DB_URL" down

echo "Running migrations..."
migrate -path db/migrations -database "$DB_URL" up

echo "Migrations applied successfully"
